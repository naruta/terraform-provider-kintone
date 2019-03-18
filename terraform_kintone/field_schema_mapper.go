package terraform_kintone

import (
	"github.com/naruta/terraform-provider-kintone/kintone"
	"github.com/naruta/terraform-provider-kintone/kintone/field"
	"github.com/pkg/errors"
)

type fieldSchemaMapper struct{}

func (m *fieldSchemaMapper) SchemaToField(fieldMap map[string]interface{}) (kintone.Field, error) {
	fieldType := kintone.FieldType(fieldMap["type"].(string))
	code := kintone.FieldCode(fieldMap["code"].(string))
	label := fieldMap["label"].(string)
	switch fieldType {
	case "SINGLE_LINE_TEXT":
		return field.NewSingleLineText(code, label), nil
	case "MULTI_LINE_TEXT":
		return field.NewMultiLineText(code, label), nil
	case "NUMBER":
		return field.NewNumber(code, label), nil
	default:
		return nil, errors.Errorf("unknown field type: %s", fieldType)
	}
}

func (m *fieldSchemaMapper) FieldToSchema(f kintone.Field) map[string]interface{} {
	return map[string]interface{}{
		"code":  f.Code().String(),
		"label": f.Label(),
		"type":  f.Type().String(),
	}
}
