package objects

type ComeCodeObject struct {
	Names  []string
	Types  []*ComeTypeObject
	Consts []ComeObjecter  // 里边存储了不会变类型的值
	Ops    [][]int  // []int 表示一条指令，[]int[0]表示指令代码，其后表示参数
}

func (co *ComeCodeObject) AddName(name string, typ *ComeTypeObject) int {
	co.Names = append(co.Names, name)
	co.Types = append(co.Types, typ)
	return len(co.Names) - 1
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

func (co *ComeCodeObject) AddConst(cons ComeObjecter) int {
	co.Consts = append(co.Consts, cons)
	return len(co.Consts) - 1
}

func (co *ComeCodeObject) AddOp(op...int) {
	co.Ops = append(co.Ops, op)
}

type ComeOpType int

const (
	ComeOp_LoadConst = iota
	ComeOp_StoreName
	ComeOp_LoadName

	ComeOp_Add
	ComeOp_Sub
	ComeOp_Multi
	ComeOp_Divide

	ComeOp_Print
	ComeOp_PrintNewLine
)

var (
	ComeComputationStringToOp = map[string]int{
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
		Ops:    [][]int{},
	}
}
