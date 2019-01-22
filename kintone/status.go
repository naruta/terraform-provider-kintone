package kintone

type StateName string

func (sn StateName) String() string {
	return string(sn)
}

type State struct {
	Name  StateName
	Index int
}

type Action struct {
	Name string
	From StateName
	To   StateName
}

type Status struct {
	Enable  bool
	States  []State
	Actions []Action
}
