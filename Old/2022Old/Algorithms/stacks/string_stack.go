package main

import "errors"

type sStack struct {
	Elements []string
	Max      int
}

func (s *sStack) Push(value string) error {
	if len(s.Elements) == s.Max && s.Max != 0 {
		return errors.New("Stack is already full")
	}

	s.Elements = append(s.Elements, value)
	return nil
}

func (s *sStack) Pop() string {
	if len(s.Elements) == 0 {
		return ""
	}

	value := s.Elements[len(s.Elements)-1]

	s.Elements = s.Elements[0 : len(s.Elements)-1]

	return value
}

func (s *sStack) Peek() string {
	if len(s.Elements) == 0 {
		return ""
	}

	return s.Elements[len(s.Elements)-1]
}

func (s *sStack) IsEmpty() bool {
	return len(s.Elements) == 0
}

func (s *sStack) IsFull() bool {
	if s.Max == 0 {
		return false
	}

	return s.Max == len(s.Elements)
}
