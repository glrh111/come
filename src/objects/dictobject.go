package objects

type ComeDictObject struct {
	ComeObject
	keyType, valueType *ComeTypeObject
	value map[ComeObjecter]ComeObjecter
}

func (o ComeDictObject) Type() *ComeTypeObject {
	return ComeDictType
}

// 变长类型
var ComeDictType = &ComeTypeObject{
	ComeObject: ComeObject{Type: ComeTypeType},
	Name: ComeDictTypeString,
	Size: 4,
}

func NewComeDictObject() *ComeDictObject {
	return &ComeDictObject{
		ComeObject: ComeObject{Type: ComeDictType},
		value: make(map[ComeObjecter]ComeObjecter),
	}
}

func (dict *ComeDictObject) Get(key ComeObjecter) (ComeObjecter, bool) {
	v, ok := dict.value[key]
	return v, ok
}

func (dict *ComeDictObject) Put(key ComeObjecter, value ComeObjecter) {
	dict.value[key] = value
}
