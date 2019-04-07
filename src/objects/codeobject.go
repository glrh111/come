package objects

type ComeCodeObject struct {
	Names  []string
	Types  []*ComeTypeObject
	Consts []ComeObjecter  // 里边存储了不会变类型的值
	Ops    []ComeOpType  // 指令
}

func (co *ComeCodeObject) AddName(name string, typ *ComeTypeObject) {
	co.Names = append(co.Names, name)
	co.Types = append(co.Types, typ)
}

func (co *ComeCodeObject) GetTypeByIndex(index int) *ComeTypeObject {
	return co.Types[index]
}

func (co *ComeCodeObject) FindNameIndex(name string) int {
	for index, value := range co.Names {
		if value == name {
			return index
		}
	}
	return -1
}

func (co *ComeCodeObject) FindConstIndex(o ComeObjecter) int {
	for index, value := range co.Consts {
		if ComeObjectEqualCompare(o, value) {
			return index
		}
	}
	return -1
}

func (co *ComeCodeObject) AddConst(cons ComeObjecter) {
	co.Consts = append(co.Consts, cons)
}

func (co *ComeCodeObject) AddOp(op ComeOpType) {
	co.Ops = append(co.Ops, op)
}

type ComeOpType int

const (
	ComeOp_LoadConst ComeOpType = iota
	ComeOp_StoreName
	ComeOp_LoadName

	ComeOp_Add
	ComeOp_Sub
	ComeOp_Multi
	ComeOp_Divide
)

var (
	ComeComputationStringToOp = map[string]ComeOpType{
		"+": ComeOp_Add,
		"-": ComeOp_Sub,
		"*": ComeOp_Multi,
		"/": ComeOp_Divide,
	}
)

func NewComeCodeObject() *ComeCodeObject {
	return &ComeCodeObject{
		Names:  []string{},
		Types:  []*ComeTypeObject{},
		Consts: []ComeObjecter{},
		Ops:    []ComeOpType{},
	}
}
