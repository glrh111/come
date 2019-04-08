package objects

import "fmt"

type ComeObjecter interface {
	Type() *ComeTypeObject
	fmt.Stringer
}


type ComeObject struct {
	Type *ComeTypeObject
}

type ComeVarObject struct {
	ComeObject
	Len int  // item amount
}

func ComeObjectEqualCompare(a ComeObjecter, b ComeObjecter) bool {
	switch v := a.(type) {
	case *ComeIntObject:
		if v2, ok := b.(*ComeIntObject); ok {
			return v.value == v2.value
		}
	case *ComeStringObject:
		if v2, ok := b.(*ComeStringObject); ok {
			return v.value == v2.value
		}
	}
	return false
}
