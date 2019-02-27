package kintone

type AppId string

func (id AppId) String() string {
	return string(id)
}

type Revision string

func (r Revision) String() string {
	return string(r)
}

type Theme string

const (
	ApplicationThemeWhite = "WHITE"
	ApplicationThemeRed   = "RED"
	ApplicationThemeBlue  = "BLUE"
	ApplicationThemeGreen = "GREEN"
	ApplicationThemeBlack = "BLACK"
)

func (t Theme) String() string {
	return string(t)
}

type Application struct {
	Id       AppId
	Setting  Setting
	Fields   []Field
	Views    []View
	Status   Status
	Revision Revision
}

type Setting struct {
	Name        string
	Description string
	Theme       Theme
}
