package objects

import (
	"fmt"
	"strings"
)

type ComeFrameObject struct {
	OpIndex int
	Stack   *ComeObjecterStack
	CO      *ComeCodeObject
	Local   map[string]ComeObjecter // 名字空间
}

func NewComeFrameObject(co *ComeCodeObject) *ComeFrameObject {
	return &ComeFrameObject{
		OpIndex: 0,
		Stack:   NewComeObjecterStack(),
		CO:      co,
		Local:   map[string]ComeObjecter{},
	}
}

// head                           at least one elem
// |---| --> |---| --> |---| --> |zero Stack|
type ComeObjecterStackNode struct {
	value ComeObjecter
	next *ComeObjecterStackNode
}

type ComeObjecterStack struct {
	head *ComeObjecterStackNode
	size int
}

func NewComeObjecterStack() *ComeObjecterStack {
	return &ComeObjecterStack{
		head: nil,
		size: 0,
	}
}

func (s *ComeObjecterStack) Pop() (ComeObjecter, bool) {
	if s.Len() == 0 {
		return nil, false
	}
	c := s.head
	s.head = s.head.next
	s.size--
	return c.value, true
}

func (s *ComeObjecterStack) Push(value ComeObjecter) {
	sn := &ComeObjecterStackNode{
		value: value,
		next: s.head,
	}
	s.head = sn
	s.size++
}

func (s *ComeObjecterStack) Len() int {
	return s.size
}

func (s *ComeObjecterStack) Peek() (ComeObjecter, bool) {
	if s.Len() == 0 {
		return nil, false
	}
	return s.head.value, true
}

func (s *ComeObjecterStack) String() string {
	sl := []string{fmt.Sprintf("Stack len-[%v]\n", s.Len())}
	sc := s.head
	for ; sc != nil; sc = sc.next {
		sl = append(sl, fmt.Sprintf("[%v]--->\n", sc.value))
	}
	return strings.Join(sl, "")
}

