package compiler

import (
	"testing"
	"fmt"
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
