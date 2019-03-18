package field

import "github.com/naruta/terraform-provider-kintone/kintone"

type MultiLineText struct {
	code  kintone.FieldCode
	label string
}

func NewMultiLineText(code kintone.FieldCode, label string) *MultiLineText {
	return &MultiLineText{
		code:  code,
		label: label,
	}
}

func (s *MultiLineText) Type() kintone.FieldType {
	return "MULTI_LINE_TEXT"
}

func (s *MultiLineText) Code() kintone.FieldCode {
	return s.code
}

func (s *MultiLineText) Label() string {
	return s.label
}
