package client

import (
	"context"
	"github.com/naruta/terraform-provider-kintone/kintone"
	"github.com/naruta/terraform-provider-kintone/kintone/raw_client"
	"strconv"
)

type ApiClientImpl struct {
	rawClient *raw_client.ApiClient
}

func New(config kintone.ApiClientConfig) kintone.ApiClient {
	rawClient := raw_client.New(raw_client.ApiClientConfig{
		Host:     config.Host,
		User:     config.User,
		Password: config.Password,
	})
	return &ApiClientImpl{
		rawClient: rawClient,
	}
}

func (c *ApiClientImpl) CreatePreviewApplication(ctx context.Context, name string) (kintone.AppId, kintone.Revision, error) {
	resp, err := raw_client.PostPreviewApp(ctx, c.rawClient, raw_client.PostPreviewAppRequest{
		Name: name,
	})
	if err != nil {
		return kintone.AppId(""), kintone.Revision(""), err
	}
	return kintone.AppId(resp.App), kintone.Revision(resp.Revision), nil
}

func (c *ApiClientImpl) UpdatePreviewApplicationSettings(ctx context.Context, appId kintone.AppId, revision kintone.Revision, setting kintone.Setting) (kintone.Revision, error) {
	resp, err := raw_client.PutPreviewAppSettings(ctx, c.rawClient, raw_client.PutPreviewAppSettingsRequest{
		App:         appId.String(),
		Revision:    revision.String(),
		Name:        setting.Name,
		Description: setting.Description,
		Theme:       setting.Theme.String(),
	})
	if err != nil {
		return kintone.Revision(""), err
	}
	return kintone.Revision(resp.Revision), nil
}

func (c *ApiClientImpl) UpdatePreviewApplicationStatus(ctx context.Context, appId kintone.AppId, revision kintone.Revision, status kintone.Status) (kintone.Revision, error) {
	states := map[string]raw_client.PutPreviewAppStatusRequestState{}
	for _, state := range status.States {
		states[state.Name.String()] = raw_client.PutPreviewAppStatusRequestState{
			Name:  state.Name.String(),
			Index: strconv.Itoa(state.Index),
		}
	}

	var actions []raw_client.PutPreviewAppStatusRequestAction
	for _, action := range status.Actions {
		actions = append(actions, raw_client.PutPreviewAppStatusRequestAction{
			Name: action.Name,
			From: action.From.String(),
			To:   action.To.String(),
		})
	}

	resp, err := raw_client.PutPreviewAppStatus(ctx, c.rawClient, raw_client.PutPreviewAppStatusRequest{
		App:      appId.String(),
		Revision: revision.String(),
		Enable:   status.Enable,
		States:   states,
		Actions:  actions,
	})
	if err != nil {
		return kintone.Revision(""), err
	}
	return kintone.Revision(resp.Revision), nil
}

func (c *ApiClientImpl) DeployApplication(ctx context.Context, appId kintone.AppId, revision kintone.Revision) error {
	_, err := raw_client.PostPreviewAppDeploy(ctx, c.rawClient, raw_client.PostPreviewAppDeployRequest{
		Apps: []raw_client.PostPreviewAppDeployRequestApp{
			{
				App:      appId.String(),
				Revision: revision.String(),
			},
		},
	})
	if err != nil {
		return err
	}

	if err := waitForDeploy(ctx, c.rawClient, appId); err != nil {
		return err
	}

	return nil
}

func (c *ApiClientImpl) DeleteApplication(ctx context.Context, appId kintone.AppId) error {
	_, err := raw_client.PostDevAppDelete(ctx, c.rawClient, raw_client.PostDevAppDeleteRequest{
		App: appId.String(),
	})
	return err
}

func (c *ApiClientImpl) FetchApplication(ctx context.Context, appId kintone.AppId) (kintone.Application, error) {
	settingsResp, err := raw_client.GetAppSettings(ctx, c.rawClient, raw_client.GetAppSettingsRequest{App: appId.String()})
	if err != nil {
		return kintone.Application{}, err
	}

	fieldsResp, err := raw_client.GetAppFormFields(ctx, c.rawClient, raw_client.GetAppFormFieldsRequest{App: appId.String()})
	if err != nil {
		return kintone.Application{}, err
	}

	statusResp, err := raw_client.GetAppStatus(ctx, c.rawClient, raw_client.GetAppStatusRequest{App: appId.String()})
	if err != nil {
		return kintone.Application{}, err
	}

	var fields []kintone.Field
	mapper := fieldPropertyMapper{}
	for _, p := range fieldsResp.Properties {
		field, err := mapper.PropertyToField(&p)
		if err != nil {
			return kintone.Application{}, err
		}
		fields = append(fields, field)
	}

	var states []kintone.State
	for _, s := range statusResp.States {
		index, err := strconv.Atoi(s.Index)
		if err != nil {
			return kintone.Application{}, err
		}
		states = append(states, kintone.State{
			Name:  kintone.StateName(s.Name),
			Index: index,
		})
	}

	var actions []kintone.Action
	for _, a := range statusResp.Actions {
		actions = append(actions, kintone.Action{
			Name: a.Name,
			From: kintone.StateName(a.From),
			To:   kintone.StateName(a.To),
		})
	}

	return kintone.Application{
		Id:       appId,
		Revision: kintone.Revision(statusResp.Revision),
		Setting: kintone.Setting{
			Name:        settingsResp.Name,
			Description: settingsResp.Description,
			Theme:       kintone.Theme(settingsResp.Theme),
		},
		Fields: fields,
		Status: kintone.Status{
			Enable:  statusResp.Enable,
			States:  states,
			Actions: actions,
		},
	}, nil
}

func (c *ApiClientImpl) CreatePreviewApplicationFormFields(ctx context.Context, appId kintone.AppId, revision kintone.Revision, fields []kintone.Field) (kintone.Revision, error) {
	properties := map[string]raw_client.PostPreviewAppFormFieldsRequestProperty{}
	mapper := fieldPropertyMapper{}
	for _, field := range fields {
		properties[field.Code().String()] = raw_client.PostPreviewAppFormFieldsRequestProperty{
			FieldProperty: mapper.FieldToProperty(field),
		}
	}

	resp, err := raw_client.PostPreviewAppFormFields(ctx, c.rawClient, raw_client.PostPreviewAppFormFieldsRequest{
		App:        appId.String(),
		Revision:   revision.String(),
		Properties: properties,
	})
	if err != nil {
		return kintone.Revision(""), err
	}
	return kintone.Revision(resp.Revision), nil
}

func (c *ApiClientImpl) UpdatePreviewApplicationFormFields(ctx context.Context, appId kintone.AppId, revision kintone.Revision, fields []kintone.Field) (kintone.Revision, error) {
	properties := map[string]raw_client.PutPreviewAppFormFieldsRequestProperty{}
	mapper := fieldPropertyMapper{}
	for _, field := range fields {
		properties[field.Code().String()] = raw_client.PutPreviewAppFormFieldsRequestProperty{
			FieldProperty: mapper.FieldToProperty(field),
		}
	}

	resp, err := raw_client.PutPreviewAppFormFields(ctx, c.rawClient, raw_client.PutPreviewAppFormFieldsRequest{
		App:        appId.String(),
		Revision:   revision.String(),
		Properties: properties,
	})
	if err != nil {
		return kintone.Revision(""), err
	}
	return kintone.Revision(resp.Revision), nil
}

func (c *ApiClientImpl) DeletePreviewApplicationFormFields(ctx context.Context, appId kintone.AppId, revision kintone.Revision, fieldCodes []kintone.FieldCode) (kintone.Revision, error) {
	var deleteFields []string
	for _, code := range fieldCodes {
		deleteFields = append(deleteFields, code.String())
	}
	resp, err := raw_client.DeletePreviewAppFormFields(ctx, c.rawClient, raw_client.DeletePreviewAppFormFieldsRequest{
		App:      appId.String(),
		Revision: revision.String(),
		Fields:   deleteFields,
	})
	if err != nil {
		return kintone.Revision(""), err
	}
	return kintone.Revision(resp.Revision), nil
}

func (c *ApiClientImpl) CreateRecord(ctx context.Context, appId kintone.AppId, record kintone.Record) (kintone.RecordId, kintone.RecordRevision, error) {
	recordValues := map[string]raw_client.PostRecordRequestRecord{}
	for key, value := range record.Values {
		recordValues[key.String()] = raw_client.PostRecordRequestRecord{
			Value: value,
		}
	}

	resp, err := raw_client.PostRecord(ctx, c.rawClient, raw_client.PostRecordRequest{
		App:    appId.String(),
		Record: recordValues,
	})
	if err != nil {
		return kintone.RecordId(""), kintone.RecordRevision(""), err
	}
	return kintone.RecordId(resp.Id), kintone.RecordRevision(resp.Revision), nil
}

func (c *ApiClientImpl) UpdateRecord(ctx context.Context, appId kintone.AppId, record kintone.Record) (kintone.RecordRevision, error) {
	recordValues := map[string]raw_client.PutRecordRequestRecord{}
	for key, value := range record.Values {
		recordValues[key.String()] = raw_client.PutRecordRequestRecord{
			Value: value,
		}
	}

	resp, err := raw_client.PutRecord(ctx, c.rawClient, raw_client.PutRecordRequest{
		App:    appId.String(),
		Id:     record.Id.String(),
		Record: recordValues,
	})
	if err != nil {
		return kintone.RecordRevision(""), err
	}
	return kintone.RecordRevision(resp.Revision), nil
}

func (c *ApiClientImpl) DeleteRecords(ctx context.Context, appId kintone.AppId, ids []kintone.RecordId) error {
	var deleteIds []string
	for _, id := range ids {
		deleteIds = append(deleteIds, id.String())
	}

	_, err := raw_client.DeleteRecords(ctx, c.rawClient, raw_client.DeleteRecordsRequest{
		App: appId.String(),
		Ids: deleteIds,
	})
	if err != nil {
		return err
	}
	return nil
}

func waitForDeploy(ctx context.Context, apiClient *raw_client.ApiClient, appId kintone.AppId) error {
	return WaitUntil(10, func() (bool, error) {
		getPreviewAppDeployResp, err := raw_client.GetPreviewAppDeploy(ctx, apiClient, raw_client.GetPreviewAppDeployRequest{
			Apps: []string{
				appId.String(),
			},
		})
		if err != nil {
			return true, err
		}

		if getPreviewAppDeployResp.Apps[0].App == appId.String() &&
			getPreviewAppDeployResp.Apps[0].Status == "SUCCESS" {
			return true, nil
		}
		return false, nil
	})
}
