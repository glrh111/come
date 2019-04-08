package objects

import "fmt"

type ComeIntObject struct {
	ComeObject
	value int
}

func (o ComeIntObject) Type() *ComeTypeObject {
	return ComeIntType
}

func (o ComeIntObject) String() string {
	return fmt.Sprintf("%v", o.value)
}

func (o *ComeIntObject) Compute(o2 *ComeIntObject, op int) *ComeIntObject {
	re := NewComeIntObject(0)
	switch op {
	case ComeOp_Add:
		re.value = o.value + o2.value
	case ComeOp_Sub:
		re.value = o.value - o2.value
	case ComeOp_Multi:
		re.value = o.value * o2.value
	case ComeOp_Divide:
		re.value = o.value / o2.value
	default:
		panic("error int compute")
	}
	return re
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