package vm

import (
	"objects"
	"fmt"
)

func Run(co *objects.ComeCodeObject) {
	// 构建 ComeFrameObject
	cf := objects.NewComeFrameObject(co)
	// run
	for cf.OpIndex = 0; cf.OpIndex < len(cf.CO.Ops); cf.OpIndex++ {
		opSlice := cf.CO.Ops[cf.OpIndex]
		switch op := opSlice[0]; op {
		case objects.ComeOp_LoadConst:
			cons := co.Consts[opSlice[1]]
			cf.Stack.Push(cons)
		case objects.ComeOp_LoadName:
			varName := co.Names[opSlice[1]]
			value := cf.Local[varName]
			cf.Stack.Push(value)
		case objects.ComeOp_StoreName:
			varName := co.Names[opSlice[1]]
			value, _ := cf.Stack.Pop()
			cf.Local[varName] = value
		case objects.ComeOp_Add, objects.ComeOp_Sub, objects.ComeOp_Multi, objects.ComeOp_Divide:
			v2, _ := cf.Stack.Pop()
			v1, _ := cf.Stack.Pop()
			v1i := v1.(*objects.ComeIntObject)
			v2i := v2.(*objects.ComeIntObject)
			v3i := v1i.Compute(v2i, op)
			cf.Stack.Push(v3i)
		case objects.ComeOp_Print:
			v, _ := cf.Stack.Pop()
			fmt.Print(v)
		case objects.ComeOp_PrintNewLine:
			fmt.Println()
		}
	}
}
