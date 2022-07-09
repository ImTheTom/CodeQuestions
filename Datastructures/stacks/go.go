package stacks

import (
	"errors"
	"fmt"
	"time"
)

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

func main() {
	start := time.Now()

	fmt.Println("Running main")
	runLimit()

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}

func runNoLimit() {
	s := &Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	fmt.Println(s.IsEmpty())

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	fmt.Println(s.IsEmpty())

	fmt.Println(s.IsFull())

	fmt.Println(s.Peek())
	fmt.Println(s.Peek())
}

func runLimit() {
	s := &Stack{
		Max: 3,
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	fmt.Println(s.IsEmpty())

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	fmt.Println(s.IsEmpty())

	fmt.Println(s.IsFull())

	fmt.Println(s.Peek())
	fmt.Println(s.Peek())
}
