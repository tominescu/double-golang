package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
	s.Push(0)
	s.Push(1)
	s.Push(3)
	s.Push(5)
	for s.Len() > 0 {
		t.Log(s.Pop().Value)
	}
}
