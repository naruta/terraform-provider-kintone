package client

import (
	"github.com/naruta/terraform-provider-kintone/kintone"
	"github.com/naruta/terraform-provider-kintone/kintone/field"
	"github.com/naruta/terraform-provider-kintone/kintone/raw_client"
	"github.com/pkg/errors"
)

type fieldPropertyMapper struct{}

func (m *fieldPropertyMapper) PropertyToField(p *raw_client.FieldProperty) (kintone.Field, error) {
	switch p.Type {
	case "SINGLE_LINE_TEXT":
		return field.NewSingleLineText(kintone.FieldCode(p.Code), p.Label), nil
	case "MULTI_LINE_TEXT":
		return field.NewMultiLineText(kintone.FieldCode(p.Code), p.Label), nil
	case "NUMBER":
		return field.NewNumber(kintone.FieldCode(p.Code), p.Label), nil
	default:
		return nil, errors.Errorf("unknown field type: %s", p.Type)
	}
}

func (m *fieldPropertyMapper) FieldToProperty(f kintone.Field) raw_client.FieldProperty {
	return raw_client.FieldProperty{
		Code:  f.Code().String(),
		Label: f.Label(),
		Type:  f.Type().String(),
	}
}
