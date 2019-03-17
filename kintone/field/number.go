package field

import "github.com/naruta/terraform-provider-kintone/kintone"

type Number struct {
	code  kintone.FieldCode
	label string
}

func NewNumber(code kintone.FieldCode, label string) *Number {
	return &Number{
		code:  code,
		label: label,
	}
}

func (s *Number) Type() kintone.FieldType {
	return "NUMBER"
}

func (s *Number) Code() kintone.FieldCode {
	return s.code
}

func (s *Number) Label() string {
	return s.label
}
