package cmd

type Stack struct {
	data []rune
	size int
}

func newStack() *Stack {
	stack := make([]rune, 0)
	return &Stack{data: stack, size: 0}
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}

func (s *Stack) push(input rune) {
	s.data = append(s.data, input)
	s.size++
	return
}

func (s *Stack) pop() rune {
	if s.isEmpty() {
		return 0
	}
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.size--
	return last
}

func (s *Stack) last() rune {
	if s.isEmpty() {
		return 0
	}
	return s.data[len(s.data)-1]
}
