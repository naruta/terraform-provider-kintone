package terraform_kintone

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/naruta/terraform-provider-kintone/kintone"
	"github.com/naruta/terraform-provider-kintone/kintone/raw_client"
)

func resourceKintoneRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceKintoneRecordCreate,
		Read:   resourceKintoneRecordRead,
		Update: resourceKintoneRecordUpdate,
		Delete: resourceKintoneRecordDelete,

		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"revision": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"values": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateKintoneRecordValues,
			},
		},
	}
}

func validateKintoneRecordValues(v interface{}, k string) (ws []string, errors []error) {
	var values map[string]string
	rawValues := v.(string)
	err := raw_client.DecodeJson([]byte(rawValues), &values)
	if err != nil {
		errors = append(errors, fmt.Errorf("invalid format of %s: %s", k, err))
	}
	return
}

func convertRecordValues(rawValues string) (map[kintone.FieldCode]string, error) {
	var valueMap map[string]string
	err := raw_client.DecodeJson([]byte(rawValues), &valueMap)
	if err != nil {
		return nil, err
	}

	values := map[kintone.FieldCode]string{}
	for key, value := range valueMap {
		values[kintone.FieldCode(key)] = value
	}

	return values, nil
}

func resourceKintoneRecordCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	apiClient := newClient(*config)
	useCase := kintone.NewCreateRecordUseCase(apiClient)
	ctx := context.Background()

	cmd := kintone.CreateRecordUseCaseCommand{
		AppId:  kintone.AppId(d.Get("app_id").(string)),
		Record: kintone.Record{},
	}

	if v, ok := d.GetOk("values"); ok {
		values, err := convertRecordValues(v.(string))
		if err != nil {
			return err
		}
		cmd.Record.Values = values
	}

	id, revision, err := useCase.Execute(ctx, cmd)
	if err != nil {
		return err
	}

	d.SetId(id.String())
	d.Set("revision", revision.String())

	return nil
}

func resourceKintoneRecordRead(d *schema.ResourceData, m interface{}) error {
	// TODO: impl
	return nil
}

func resourceKintoneRecordUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	apiClient := newClient(*config)
	useCase := kintone.NewUpdateRecordUseCase(apiClient)
	ctx := context.Background()

	cmd := kintone.UpdateRecordUseCaseCommand{
		AppId: kintone.AppId(d.Get("app_id").(string)),
		Record: kintone.Record{
			Id:       kintone.RecordId(d.Id()),
			Revision: kintone.RecordRevision(d.Get("revision").(string)),
		},
	}

	if v, ok := d.GetOk("values"); ok {
		values, err := convertRecordValues(v.(string))
		if err != nil {
			return err
		}
		cmd.Record.Values = values
	}

	revision, err := useCase.Execute(ctx, cmd)
	if err != nil {
		return err
	}

	d.Set("revision", revision.String())

	return nil
}

func resourceKintoneRecordDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	apiClient := newClient(*config)
	useCase := kintone.NewDeleteRecordUseCase(apiClient)
	ctx := context.Background()

	cmd := kintone.DeleteRecordUseCaseCommand{
		AppId: kintone.AppId(d.Get("app_id").(string)),
		Id:    kintone.RecordId(d.Id()),
	}

	return useCase.Execute(ctx, cmd)
}
