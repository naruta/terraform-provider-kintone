package client

import (
	"github.com/naruta/terraform-provider-kintone/kintone/raw_client"
	"reflect"
	"testing"
)

func TestFieldPropertyMapper(t *testing.T) {
	testCases := []struct {
		title         string
		property      raw_client.FieldProperty
		shouldBeError bool
	}{
		{
			title: "SINGLE_LINE_TEXT",
			property: raw_client.FieldProperty{
				Type:  "SINGLE_LINE_TEXT",
				Code:  "text-1",
				Label: "🍣🍺",
			},
		},
		{
			title: "MULTI_LINE_TEXT",
			property: raw_client.FieldProperty{
				Type:  "MULTI_LINE_TEXT",
				Code:  "text-2",
				Label: "🍣🍺",
			},
		},
		{
			title: "NUMBER",
			property: raw_client.FieldProperty{
				Type:  "NUMBER",
				Code:  "number-1",
				Label: "🍣🍺",
			},
		},
		{
			title: "Unknown type",
			property: raw_client.FieldProperty{
				Type:  "ABCDEFG",
				Code:  "xxx-1",
				Label: "🍣🍺",
			},
			shouldBeError: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.title, func(t *testing.T) {
			mapper := fieldPropertyMapper{}
			f, err := mapper.PropertyToField(&tt.property)
			if tt.shouldBeError {
				if err == nil {
					t.Fatalf("expected: error, actual: no errors")
				}
				return
			}
			if err != nil {
				t.Fatalf("error: %+v", err)
			}

			property := mapper.FieldToProperty(f)
			if !reflect.DeepEqual(property, tt.property) {
				t.Fatalf("property != tt.property: property=%+v, tt.property=%+v", property, tt.property)
			}
		})
	}
}
