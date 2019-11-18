package kintone

type ViewType string

const (
	ViewTypeList     = ViewType("LIST")
	ViewTypeCalendar = ViewType("CALENDAR")
	ViewTypeCustom   = ViewType("CUSTOM")
)

func (vt ViewType) String() string {
	return string(vt)
}

type View struct {
	Index       string
	Type        ViewType
	BuiltinType string
	Name        string
	Fields      []FieldCode
}
