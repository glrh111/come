package compiler

import (
	"io"
	"objects"
	"io/ioutil"
	"strings"
	"regexp"
	"fmt"
	"strconv"
)

const (
	LineSeparator = "\n"
	CommentSign = "--"


)

var (
	// regex match
	MatchRegex_Declaration = regexp.MustCompile(
		fmt.Sprintf(`(?P<typeName>[%v|%v|%v|%v])\s*(?P<varName>\S+)\s*;`, objects.ComeIntType, objects.ComeStringType, objects.ComeDictType),
	)
	MatchRegex_Assignment = regexp.MustCompile(`(?P<varName>\S+)\s*=(?P<varValue>\S+)\s*;`)
	MatchRegex_Compute = regexp.MustCompile(`(?P<varName>\S+)\s*=\s*(?P<varName1>\S+)\s*(?P<op>[+|-|*|/])\s*(?P<varName2>\S+)\s;`)
	MatchRegex_Print = regexp.MustCompile(`print\(\s*(?P<varName>\S+)\s*\)\s*;`)
)

var testSrcCode = `
-- testSrcCode 注释 先声明，再赋值
int a;
a = 5;
int b;
b = 6;
int c;
c = a + b;
print(c);
`

type matchDeclarationResult struct {
	Type *objects.ComeTypeObject
	VarName string
}

type matchAssignmentResult struct {
	VarName string
	ValueInString string
}

type matchComputationResult struct {
	leftVarName string
	rightVarName1 string
	rightVarName2 string
	op objects.ComeOpType  // op
}

type matchPrintResult struct {
	VarName string
}

// int b
func matchDeclaration(line string) (*matchDeclarationResult, bool) {
	strs := strings.Fields(line);
	if len(strs) != 2 {
		return nil, false
	}

	if t, ok := objects.ComeTypeStringToType[strs[0]]; ok { // find type
		return &matchDeclarationResult{
			Type: t,
			VarName: strs[1],
		}, true
	} else {
		return nil, false
	}
}

// a = 1
func matchAssignment(line string) (*matchAssignmentResult, bool) {
	strs := strings.Fields(line);
	if len(strs) != 3 && strs[1] != "=" {
		return nil, false
	}

	return &matchAssignmentResult{
		VarName: strs[0],
		ValueInString: strs[2],
	}, true
}

// 0 1 2 3 4
// c = a + b
func matchComputation(line string) (*matchComputationResult, bool) {
	strs := strings.Fields(line);
	if len(strs) != 5 && strs[1] != "=" {
		return nil, false
	}

	if op, ok := objects.ComeComputationStringToOp[strs[3]]; ok {
		return &matchComputationResult{
			leftVarName: strs[0],
			rightVarName1: strs[2],
			rightVarName2: strs[4],
			op: op,
		}, true
	} else {
		return nil, false
	}
}

// print(wocao)
func matchPrint(line string) (*matchPrintResult, bool) {
	if !(len(line) >= 8 && line[:5] == "print") {
		return nil, false
	}
	return &matchPrintResult{
		VarName:line[6:len(line)-1],
	}, true
}

func Compiler(f io.Reader) *objects.ComeCodeObject {
	// 读取line
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic("Read Err")
	}
	var (
		lines = strings.Split(string(bs), LineSeparator)
		lineNo = 1
		co = objects.NewComeCodeObject()
	)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) >= 2 && line[:2] == CommentSign {
			continue
		}
		if len(line) >= 1 {
			if line[len(line)-1] != ';' {
				panic(
					fmt.Sprintf("line [%v] not end with [;]", lineNo),
				)
			}
		} else {
			continue
		}
		line = line[:len(line)-1] // 去除 ;

		lineNo += 1
		if md, ok := matchDeclaration(line); ok {
			// check
			nameIndex := co.FindNameIndex(md.VarName)
			if nameIndex >= 0 {
				panic(fmt.Sprintf("redeclaration line[%v] var[%v]", lineNo-1, md.VarName))
			}
			// add to
			co.AddName(md.VarName, md.Type)
			continue
		}
		if ma, ok := matchAssignment(line); ok {
			// check if declaration
			nameIndex := co.FindNameIndex(ma.VarName)
			if nameIndex < 0 {
				panic(fmt.Sprintf("assign before declaration line[%v] var[%v]", lineNo-1, ma.VarName))
			}
			var (
				value objects.ComeObjecter
				typ = co.GetTypeByIndex(nameIndex)
			)
			switch typ {
			case objects.ComeIntType:
				v, err := strconv.Atoi(ma.ValueInString)
				if err != nil {
					panic(fmt.Sprintf("not int value [%v] line [%v]", ma.ValueInString, lineNo-1))
				}
				value = objects.NewComeIntObject(v)
			case objects.ComeStringType:
				value = objects.NewComeStringObject(ma.ValueInString)
			default:
				panic(fmt.Sprintf("unknown value [%v] line [%v]", ma.ValueInString, lineNo-1))
			}
			// check value
			valueIndex := co.FindConstIndex(value)
			if valueIndex < 0 { // add to const
				co.AddConst(value)
			}
			// op code TODO 加入参数
			co.AddOp(objects.ComeOp_LoadConst)
			co.AddOp(objects.ComeOp_StoreName)
			continue
		}
		//if mc, ok := matchComputation(line); ok {
		//
		//	continue
		//}
		//if mp, ok := matchPrint(line); ok {
		//
		//	continue
		//}

		panic(
			fmt.Sprintf("unrecognized line [%v]", lineNo),
		)

	}
	return nil
}
