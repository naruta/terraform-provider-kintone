package field

import "github.com/naruta/terraform-provider-kintone/kintone"

type SingleLineText struct {
	code  kintone.FieldCode
	label string
}

func NewSingleLineText(code kintone.FieldCode, label string) *SingleLineText {
	return &SingleLineText{
		code:  code,
		label: label,
	}
}

func (s *SingleLineText) Type() kintone.FieldType {
	return "SINGLE_LINE_TEXT"
}

func (s *SingleLineText) Code() kintone.FieldCode {
	return s.code
}

func (s *SingleLineText) Label() string {
	return s.label
}
