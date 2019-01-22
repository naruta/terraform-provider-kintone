package kintone

type FieldType string

const (
	FieldSingleLineText = "SINGLE_LINE_TEXT"
	FieldNumber         = "NUMBER"
)

func (ft FieldType) String() string {
	return string(ft)
}

type FieldCode string

func (fc FieldCode) String() string {
	return string(fc)
}

type Field struct {
	Code      FieldCode
	Label     string
	FieldType FieldType
}
