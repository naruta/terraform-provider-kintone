package kintone

type RecordId string

func (rid RecordId) String() string {
	return string(rid)
}

type RecordRevision string

func (r RecordRevision) String() string {
	return string(r)
}

type Record struct {
	Id       RecordId
	Revision RecordRevision
	Values   map[FieldCode]string
}
