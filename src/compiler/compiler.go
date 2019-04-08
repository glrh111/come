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
	StringSeparator = "\""


)

var (
	// regex match
	MatchRegex_Declaration = regexp.MustCompile(
		fmt.Sprintf(`(?P<typeName>[%v|%v|%v|%v])\s*(?P<varName>\S+)\s*;`, objects.ComeIntType, objects.ComeStringType, objects.ComeDictType),
	)
	MatchRegex_Assignment = regexp.MustCompile(`(?P<varName>\S+)\s*=(?P<varValue>\S+)\s*;`)
	MatchRegex_Compute = regexp.MustCompile(`(?P<varName>\S+)\s*=\s*(?P<varName1>\S+)\s*(?P<op>[+|-|*|/])\s*(?P<varName2>\S+)\s;`)
	MatchRegex_Print = regexp.MustCompile(`print\(\s*(?P<varName>\S+)\s*\)\s*;`)

	TestSrcCode = `
-- comelang 的第一个程序
int a;
a = 5;

int b;
b = 6;

int c;
c = a + b;

print(c);

string d;
d = "hello, world!";
print(d);
`
)

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
	op int  // op
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
	if strings.Contains(line, StringSeparator) { // b = "hello, world!"
		strs := strings.Split(line, "=")
		if len(strs) != 2 {
			return nil, false
		}
		s := strings.TrimSpace(strs[1])
		s = s[1:len(s)-1]
		return &matchAssignmentResult{
			VarName: strings.TrimSpace(strs[0]),
			ValueInString: s,
		}, true
	} else { // a = 1
		strs := strings.Fields(line)
		if !(len(strs) == 3 && strs[1] == "=") {
			return nil, false
		}
		return &matchAssignmentResult{
			VarName: strs[0],
			ValueInString: strs[2],
		}, true
	}
}

// 0 1 2 3 4
// c = a + b
func matchComputation(line string) (*matchComputationResult, bool) {
	strs := strings.Fields(line);
	if !(len(strs) == 5 && strs[1] == "=") {
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

func Compile(f io.Reader) *objects.ComeCodeObject {
	// 读取line
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic("Read Err")
	}
	var (
		lines = strings.Split(string(bs), LineSeparator)
		co = objects.NewComeCodeObject()
	)
	for lineNo, line := range lines {
		lineNo += 1
		line = strings.TrimSpace(line)
		if len(line) >= 2 && line[:2] == CommentSign {
			goto innerend
		}
		if len(line) >= 1 {
			if line[len(line)-1] != ';' {
				panic(
					fmt.Sprintf("line [%v] not end with [;]", lineNo),
				)
			}
		} else {
			goto innerend
		}
		line = line[:len(line)-1] // 去除 ;


		if md, ok := matchDeclaration(line); ok {
			// check
			nameIndex := co.FindNameIndex(md.VarName)
			if nameIndex >= 0 {
				panic(fmt.Sprintf("redeclaration line[%v] var[%v]", lineNo, md.VarName))
			}
			// add to
			co.AddName(md.VarName, md.Type)
			goto innerend
		}
		// a = 1
		// LOAD_CONST const_index
		// STORE_NAME name_index
		if ma, ok := matchAssignment(line); ok {
			// check if declared
			nameIndex := co.FindNameIndex(ma.VarName)
			if nameIndex < 0 {
				panic(fmt.Sprintf("assign before declaration line[%v] var[%v]", lineNo, ma.VarName))
			}
			var (
				value objects.ComeObjecter
				typ = co.GetTypeByIndex(nameIndex)
			)
			switch typ {
			case objects.ComeIntType:
				v, err := strconv.Atoi(ma.ValueInString)
				if err != nil {
					panic(fmt.Sprintf("not int value [%v] line [%v]", ma.ValueInString, lineNo))
				}
				value = objects.NewComeIntObject(v)
			case objects.ComeStringType:
				value = objects.NewComeStringObject(ma.ValueInString)
			default:
				panic(fmt.Sprintf("unknown value [%v] line [%v]", ma.ValueInString, lineNo))
			}
			// check value
			valueIndex := co.FindConstIndex(value)
			if valueIndex < 0 { // add to const
				valueIndex = co.AddConst(value)
			}
			// op code
			co.AddOp(objects.ComeOp_LoadConst, valueIndex)
			co.AddOp(objects.ComeOp_StoreName, nameIndex)
			goto innerend
		}
		// c = a + b
		// LOAD_NAME name_a_index
		// LOAD_NAME name_b_index
		// ADD
		// STORE_NAME name_c_index
		if mc, ok := matchComputation(line); ok {
			var (
				nameIndex1 = co.FindNameIndex(mc.leftVarName)
				nameIndex2 = co.FindNameIndex(mc.rightVarName1)
				nameIndex3 = co.FindNameIndex(mc.rightVarName2)
			)

			// check if a, b, c declared
			for _, nameIndex := range []int{nameIndex1, nameIndex2, nameIndex3} {
				if nameIndex < 0 {
					panic(fmt.Sprintf("undeclared line [%v]", lineNo))
				}
			}
			// load name
			co.AddOp(objects.ComeOp_LoadName, nameIndex2)
			co.AddOp(objects.ComeOp_LoadName, nameIndex3)
			co.AddOp(mc.op)
			co.AddOp(objects.ComeOp_StoreName, nameIndex1)
			goto innerend
		}
		// print(a)
		// LOAD_NAME name_index
		// PRINT
		// PRINT_NEWLINE
		if mp, ok := matchPrint(line); ok {
			nameIndex := co.FindNameIndex(mp.VarName)
			if nameIndex < 0 {
				panic(fmt.Sprintf("undeclared name [%v] line [%v]", mp.VarName, lineNo))
			}
			co.AddOp(objects.ComeOp_LoadName, nameIndex)
			co.AddOp(objects.ComeOp_Print)
			co.AddOp(objects.ComeOp_PrintNewLine)
			goto innerend
		}

		panic(
			fmt.Sprintf("unrecognized line [%v]", lineNo),
		)
innerend:
	}
	return co
}
