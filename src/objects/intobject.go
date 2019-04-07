package objects

type ComeIntObject struct {
	ComeObject
	value int
}

func (o ComeIntObject) Type() *ComeTypeObject {
	return ComeIntType
}

var ComeIntType = &ComeTypeObject{
	ComeObject: ComeObject{Type: ComeTypeType},
	Name: ComeIntTypeString,
	Size: 4,
}

func NewComeIntObject(value int) *ComeIntObject {
	return &ComeIntObject{
		ComeObject: ComeObject{Type: ComeIntType},
		value: value,
	}
}