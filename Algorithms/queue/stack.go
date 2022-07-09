package main

import "errors"

type Stack struct {
	Elements []int
	Max      int
}

func (s *Stack) Push(value int) error {
	if len(s.Elements) == s.Max && s.Max != 0 {
		return errors.New("Stack is already full")
	}

	s.Elements = append(s.Elements, value)
	return nil
}

func (s *Stack) Pop() int {
	if len(s.Elements) == 0 {
		return 0
	}

	value := s.Elements[len(s.Elements)-1]

	s.Elements = s.Elements[0 : len(s.Elements)-1]

	return value
}

func (s *Stack) Peek() int {
	if len(s.Elements) == 0 {
		return 0
	}

	return s.Elements[len(s.Elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}

func (s *Stack) IsFull() bool {
	if s.Max == 0 {
		return false
	}

	return s.Max == len(s.Elements)
}
