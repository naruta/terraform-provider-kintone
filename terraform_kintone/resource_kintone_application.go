package terraform_kintone

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/naruta/terraform-provider-kintone/kintone"
	"github.com/naruta/terraform-provider-kintone/kintone/client"
	"strings"
)

func resourceKintoneApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceKintoneApplicationCreate,
		Read:   resourceKintoneApplicationRead,
		Update: resourceKintoneApplicationUpdate,
		Delete: resourceKintoneApplicationDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"theme": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					kintone.ApplicationThemeWhite,
					kintone.ApplicationThemeRed,
					kintone.ApplicationThemeBlue,
					kintone.ApplicationThemeGreen,
					kintone.ApplicationThemeBlack,
				}, false),
			},
			"revision": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"field": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:     schema.TypeString,
							Required: true,
						},
						"label": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								kintone.FieldSingleLineText,
								kintone.FieldNumber,
								kintone.FieldMultiLineText,
							}, false),
						},
					},
				},
			},
			"status_enable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"index": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
			"action": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"from": {
							Type:     schema.TypeString,
							Required: true,
						},
						"to": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func newClient(config Config) kintone.ApiClient {
	return client.New(kintone.ApiClientConfig{
		Host:     config.Host,
		User:     config.User,
		Password: config.Password,
	})
}

func convertFields(fieldSet *schema.Set) ([]kintone.Field, error) {
	var fields []kintone.Field
	mapper := fieldSchemaMapper{}
	for _, fieldMap := range fieldSet.List() {
		field, err := mapper.SchemaToField(fieldMap.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}
	return fields, nil
}

func convertState(stateSet *schema.Set) []kintone.State {
	var states []kintone.State
	for _, stateMap := range stateSet.List() {
		state := stateMap.(map[string]interface{})
		states = append(states, kintone.State{
			Name:  kintone.StateName(state["name"].(string)),
			Index: state["index"].(int),
		})
	}
	return states
}

func convertAction(actionList []interface{}) []kintone.Action {
	var actions []kintone.Action
	for _, actionMap := range actionList {
		action := actionMap.(map[string]interface{})
		actions = append(actions, kintone.Action{
			Name: action["name"].(string),
			From: kintone.StateName(action["from"].(string)),
			To:   kintone.StateName(action["to"].(string)),
		})
	}
	return actions
}

func resourceKintoneApplicationCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	apiClient := newClient(*config)
	useCase := kintone.NewCreateApplicationUseCase(apiClient)
	ctx := context.Background()

	cmd := kintone.CreateApplicationUseCaseCommand{
		Setting: kintone.Setting{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Theme:       kintone.Theme(d.Get("theme").(string)),
		},
	}

	if v, ok := d.GetOk("field"); ok {
		fieldSet := v.(*schema.Set)
		fields, err := convertFields(fieldSet)
		if err != nil {
			return err
		}
		cmd.Fields = fields
	}

	status := kintone.Status{
		Enable: d.Get("status_enable").(bool),
	}
	if v, ok := d.GetOk("state"); ok {
		stateSet := v.(*schema.Set)
		status.States = convertState(stateSet)
	}
	if v, ok := d.GetOk("action"); ok {
		actionList := v.([]interface{})
		status.Actions = convertAction(actionList)
	}
	cmd.Status = status

	appId, revision, err := useCase.Execute(ctx, cmd)
	if err != nil {
		return err
	}

	d.SetId(appId.String())
	d.Set("revision", revision.String())

	return nil
}

func resourceKintoneApplicationRead(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	apiClient := newClient(*config)
	useCase := kintone.NewFetchApplicationUseCase(apiClient)
	ctx := context.Background()

	appId := kintone.AppId(d.Id())

	app, err := useCase.Execute(ctx, kintone.FetchApplicationUseCaseQuery{AppId: appId})
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", app.Setting.Name)
	d.Set("description", app.Setting.Description)
	d.Set("theme", app.Setting.Theme)
	d.Set("revision", app.Revision.String())

	fields := make([]map[string]interface{}, 0, len(app.Fields))
	mapper := fieldSchemaMapper{}
	for _, f := range app.Fields {
		fields = append(fields, mapper.FieldToSchema(f))
	}
	d.Set("field", fields)

	d.Set("status_enable", app.Status.Enable)

	states := make([]map[string]interface{}, 0, len(app.Status.States))
	for _, s := range app.Status.States {
		states = append(states, map[string]interface{}{
			"name":  s.Name,
			"index": s.Index,
		})
	}
	d.Set("state", states)

	actions := make([]map[string]interface{}, 0, len(app.Status.Actions))
	for _, a := range app.Status.Actions {
		actions = append(actions, map[string]interface{}{
			"name": a.Name,
			"from": a.From.String(),
			"to":   a.To.String(),
		})
	}
	d.Set("action", actions)

	return nil
}

func findField(fields []kintone.Field, target kintone.Field) (kintone.Field, bool) {
	for _, field := range fields {
		if field.Code() == target.Code() {
			return field, true
		}
	}
	return nil, false
}

func diff(oldFieldSet, newFieldSet *schema.Set) ([]kintone.Field, []kintone.Field, []kintone.FieldCode, error) {
	oldFields, err := convertFields(oldFieldSet)
	if err != nil {
		return nil, nil, nil, err
	}
	newFields, err := convertFields(newFieldSet)
	if err != nil {
		return nil, nil, nil, err
	}

	var createFields []kintone.Field
	var updateFields []kintone.Field
	var deleteFieldCodes []kintone.FieldCode

	for _, newField := range newFields {
		oldField, found := findField(oldFields, newField)
		if found {
			if newField != oldField {
				updateFields = append(updateFields, newField)
				continue
			}
		} else {
			createFields = append(createFields, newField)
		}
	}

	for _, oldField := range oldFields {
		_, found := findField(newFields, oldField)
		if !found {
			deleteFieldCodes = append(deleteFieldCodes, oldField.Code())
		}
	}

	return createFields, updateFields, deleteFieldCodes, nil
}

func resourceKintoneApplicationUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	apiClient := newClient(*config)
	useCase := kintone.NewUpdateApplicationUseCase(apiClient)
	ctx := context.Background()

	appId := kintone.AppId(d.Id())
	revision := kintone.Revision(d.Get("revision").(string))

	var createFields, updateFields []kintone.Field
	var deleteFieldCodes []kintone.FieldCode
	if d.HasChange("field") {
		oldFieldSet, newFieldSet := d.GetChange("field")
		var err error
		createFields, updateFields, deleteFieldCodes, err = diff(oldFieldSet.(*schema.Set), newFieldSet.(*schema.Set))
		if err != nil {
			return err
		}
	}

	status := kintone.Status{
		Enable: d.Get("status_enable").(bool),
	}
	if v, ok := d.GetOk("state"); ok {
		stateSet := v.(*schema.Set)
		status.States = convertState(stateSet)
	}
	if v, ok := d.GetOk("action"); ok {
		actionList := v.([]interface{})
		status.Actions = convertAction(actionList)
	}

	revision, err := useCase.Execute(ctx, kintone.UpdateApplicationUseCaseCommand{
		AppId: appId,
		Setting: kintone.Setting{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Theme:       kintone.Theme(d.Get("theme").(string)),
		},
		Status:           status,
		Revision:         revision,
		CreateFields:     createFields,
		UpdateFields:     updateFields,
		DeleteFieldCodes: deleteFieldCodes,
	})
	if err != nil {
		return err
	}

	d.Set("revision", revision)

	return nil
}

func resourceKintoneApplicationDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	apiClient := newClient(*config)
	useCase := kintone.NewDeleteApplicationUseCase(apiClient)
	ctx := context.Background()

	cmd := kintone.DeleteApplicationUseCaseCommand{
		AppId: kintone.AppId(d.Id()),
	}

	return useCase.Execute(ctx, cmd)
}
