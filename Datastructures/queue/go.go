package main

import (
	"fmt"
	"time"
)

type Queue struct {
	Elements []int
}

func (q *Queue) Enqueue(value int) {
	q.Elements = append(q.Elements, value)
}

func (q *Queue) Dequeue() int {
	if len(q.Elements) == 0 {
		return 0
	}

	element := q.Elements[0]

	q.Elements = q.Elements[1:len(q.Elements)]

	return element
}

func (q *Queue) Size() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func main() {
	start := time.Now()

	fmt.Println("Running main")

	q := &Queue{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
