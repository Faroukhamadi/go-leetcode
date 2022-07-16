package main

type MinStack struct {
	elements []*MinStackElement
}

type MinStackElement struct {
	Val     int
	Minimum int
}

func Constructor() MinStack {
	return MinStack{elements: make([]*MinStackElement, 0, 50)}
}

func (s *MinStack) Push(val int) {
	newElem := &MinStackElement{Val: val}

	if len(s.elements) > 0 {
		prevMin := s.elements[len(s.elements)-1].Minimum
		if val < prevMin {
			newElem.Minimum = val
		} else {
			newElem.Minimum = prevMin
		}
	} else {
		newElem.Minimum = val
	}

	s.elements = append(s.elements, newElem)
}

func (s *MinStack) Pop() {
	s.elements = s.elements[0 : len(s.elements)-1]
}

func (s *MinStack) Top() int {
	return s.elements[len(s.elements)-1].Val
}

func (s *MinStack) GetMin() int {
	return s.elements[len(s.elements)-1].Minimum
}

func main() {
}
