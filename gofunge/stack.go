package gofunge

type Stack struct {
	stack []int
}

// pop a value from the stack, panic if the stack is empty
func (s *Stack) Pop() int {
	if len(s.stack) == 0 {
		panic("Stack is empty!")
	}
	value := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return value
}
func (s *Stack) Push(val int) int {
	s.stack = append(s.stack, val)
	return val
}

func (s *Stack) Peek() int {
	return s.stack[len(s.stack)-1]
}
