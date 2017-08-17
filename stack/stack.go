package stack

type Element struct {
	Value interface{}
}

type Stack struct {
	elems []*Element
}

func New() *Stack {
	return &Stack{elems: make([]*Element, 0)}
}

func (s *Stack) Push(e interface{}) {
	elem := &Element{Value: e}
	s.elems = append(s.elems, elem)
}

func (s *Stack) Pop() *Element {
	n := len(s.elems)
	if n == 0 {
		return nil
	}
	e := s.elems[n-1]
	s.elems = s.elems[:n-1]
	return e
}

func (s *Stack) Top() *Element {
	n := len(s.elems)
	if n == 0 {
		return nil
	}
	return s.elems[n-1]
}

func (s *Stack) Len() int {
	return len(s.elems)
}
