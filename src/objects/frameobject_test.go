package objects

import (
	"testing"
	"fmt"
)

func TestComeObjecterStack(t *testing.T) {
	// new
	stk := NewComeObjecterStack()

	// empty len
	emptyLen := stk.Len()
	if emptyLen != 0 {
		t.Error("err emptyLen")
		t.Fail()
	}
	fmt.Print(stk)

	// PUSH
	intElems := []int{1, 2, 3, 4, 5}
	for _, value := range intElems {
		stk.Push(NewComeIntObject(value))
	}

	// show
	fmt.Print(stk)

	// Peek and pop
	for i := len(intElems)-1; i >= 0; i-- {
		ex := NewComeIntObject(intElems[i])

		p, okp := stk.Peek()
		if okp != true {
			t.Error("error peek-ok")
			t.Fail()
		}
		if ComeObjectEqualCompare(ex, p) != true {
			t.Error("error peek")
			t.Fail()
		}

		po, okpo := stk.Pop()
		if okpo != true {
			t.Error("error pop-ok")
			t.Fail()
		}
		if ComeObjectEqualCompare(ex, po) != true {
			t.Error("error pop")
			t.Fail()
		}
	}

	// empty pop
	p, okp := stk.Peek()
	if okp != false {
		t.Error("error peek-ok")
		t.Fail()
	}
	if p  != nil {
		t.Error("error peek")
		t.Fail()
	}

	po, okpo := stk.Pop()
	if okpo != false {
		t.Error("error pop-ok")
		t.Fail()
	}
	if po != nil {
		t.Error("error pop")
		t.Fail()
	}

	// empty len
	emptyLen2 := stk.Len()
	if emptyLen2 != 0 {
		t.Error("err emptyLen")
		t.Fail()
	}
}
