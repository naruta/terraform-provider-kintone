package terraform_kintone

import (
	"reflect"
	"testing"
)

func TestFieldSchemaMapper(t *testing.T) {
	testCases := []struct {
		title         string
		fieldMap      map[string]interface{}
		shouldBeError bool
	}{
		{
			title: "SINGLE_LINE_TEXT",
			fieldMap: map[string]interface{}{
				"type":  "SINGLE_LINE_TEXT",
				"code":  "text-1",
				"label": "🍣🍺",
			},
		},
		{
			title: "MULTI_LINE_TEXT",
			fieldMap: map[string]interface{}{
				"type":  "MULTI_LINE_TEXT",
				"code":  "text-2",
				"label": "🍣🍺",
			},
		},
		{
			title: "NUMBER",
			fieldMap: map[string]interface{}{
				"type":  "NUMBER",
				"code":  "number-1",
				"label": "🍣🍺",
			},
		},
		{
			title: "Unknown type",
			fieldMap: map[string]interface{}{
				"type":  "ABCDEFG",
				"code":  "xxx-1",
				"label": "🍣🍺",
			},
			shouldBeError: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.title, func(t *testing.T) {
			mapper := fieldSchemaMapper{}
			f, err := mapper.SchemaToField(tt.fieldMap)
			if tt.shouldBeError {
				if err == nil {
					t.Fatalf("expected: error, actual: no errors")
				}
				return
			}
			if err != nil {
				t.Fatalf("error: %+v", err)
			}

			fieldMap := mapper.FieldToSchema(f)
			if !reflect.DeepEqual(fieldMap, tt.fieldMap) {
				t.Fatalf("fieldMap != tt.fieldMap: fieldMap=%+v, tt.fieldMap=%+v", fieldMap, tt.fieldMap)
			}
		})
	}
}
