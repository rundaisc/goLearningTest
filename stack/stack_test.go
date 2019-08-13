package stack

import "testing"

func TestStack_Push(t *testing.T) {
	s := Stack{}
	for index := 1; index < 11; index++ {
		s.Push(index)
	}
	t.Log(s)
	for index := 1; index < 11; index++ {
		t.Log(s.Pop())

	}
}


