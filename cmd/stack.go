package cmd

type Stack struct {
	data []rune
}

func newStack() *Stack {
	stack := make([]rune, 0)
	return &Stack{data: stack}
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) push(input rune) {
	s.data = append(s.data, input)
}

func (s *Stack) pop() rune {
	if s.isEmpty() {
		return 0
	}
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return last
}

func (s *Stack) last() rune {
	if s.isEmpty() {
		return 0
	}
	return s.data[len(s.data)-1]
}
