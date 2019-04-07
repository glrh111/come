package objects

const (
	ComeIntTypeString    = "int"
	ComeStringTypeString = "string"
	ComeDictTypeString   = "dict"
)

// object 的类型
type ComeTypeObject struct {
	ComeObject
	Name string
	Size int // byte
}

var (
	ComeTypeType         *ComeTypeObject
	ComeTypeStringToType = map[string]*ComeTypeObject{
		ComeIntTypeString:    ComeIntType,
		ComeStringTypeString: ComeStringType,
		ComeDictTypeString:   ComeDictType,
	}
)

func init() {
	ComeTypeType = &ComeTypeObject{
		//ComeObject{Type: ComeTypeType},
		Name: "type",
		//Size: unsafe.Sizeof(&ComeTypeObject{}),
	}
	ComeTypeType.Type = ComeTypeType
}
