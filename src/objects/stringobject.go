package objects

type ComeStringObject struct {
	ComeObject
	value string
}

func (o ComeStringObject) Type() *ComeTypeObject {
	return ComeStringType
}

func (o ComeStringObject) String() string {
	return o.value
}

// 变长类型
var ComeStringType = &ComeTypeObject{
	ComeObject: ComeObject{Type: ComeTypeType},
	Name: ComeStringTypeString,
	Size: 4,
}

func NewComeStringObject(value string) *ComeStringObject {
	return &ComeStringObject{
		ComeObject: ComeObject{Type: ComeStringType},
		value: value,
	}
}