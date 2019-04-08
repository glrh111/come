package compiler

import (
	"testing"
	"fmt"
	"strings"
)

func TestComeCodeObjectRegex(t *testing.T) {
	var (
		declarationStr = "int a;"
		//declarationNotStr = "cao b;"
		//assignmentStr = "a= 7;"
		//assignmentNotStr = "b==9;"
	)
	// 声明
	groupNames := MatchRegex_Declaration.SubexpNames()
	matchRe := MatchRegex_Declaration.FindAllStringSubmatch(declarationStr, -1)
	fmt.Println(len(groupNames), groupNames) // "" "typeName" "varName"
	fmt.Println(matchRe)

	// 赋值

}

func TestComeCodeObject_matchDeclaration(t *testing.T) {
	var (
		s = "int a"
		ns = "cao b"
	)
	sr, ok1 := matchDeclaration(s)
	nsr, ok2 := matchDeclaration(ns)

	fmt.Printf("%+v\n%+v\n", sr, nsr)
	if ok1 != true {
		t.Error("error ok1")
		t.Fail()
	}
	if ok2 != false {
		t.Error("error ok2")
		t.Fail()
	}
}

func TestComeCodeObject_matchAssignment(t *testing.T) {
	var (
		s = "a = 6"
		ns = "b != 7"
	)
	sr, ok1 := matchAssignment(s)
	nsr, ok2 := matchAssignment(ns)

	fmt.Printf("%+v\n%+v\n", sr, nsr)
	if ok1 != true {
		t.Error("error ok1")
		t.Fail()
	}
	if ok2 != false {
		t.Error("error ok2")
		t.Fail()
	}
}

func TestComeCodeObject_matchComputation(t *testing.T) {
	var (
		s = "a = 7 + 8"
		ns = "b =9"
	)
	sr, ok1 := matchComputation(s)
	nsr, ok2 := matchComputation(ns)

	fmt.Printf("%+v\n%+v\n", sr, nsr)
	if ok1 != true {
		t.Error("error ok1")
		t.Fail()
	}
	if ok2 != false {
		t.Error("error ok2")
		t.Fail()
	}
}

func TestComeCodeObject_matchPrint(t *testing.T) {
	var (
		s = "print(a)"
		ns = "cao(b)"
	)
	sr, ok1 := matchPrint(s)
	nsr, ok2 := matchPrint(ns)

	fmt.Printf("%+v\n%+v\n", sr, nsr)
	if ok1 != true {
		t.Error("error ok1")
		t.Fail()
	}
	if ok2 != false {
		t.Error("error ok2")
		t.Fail()
	}
}

func TestCompile(t *testing.T) {
	var (
		demo = TestSrcCode
	)
	co := Compile(strings.NewReader(demo))
	fmt.Printf("%+v\n", co)
}