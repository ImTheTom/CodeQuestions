package main

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

func (q *Queue) Empty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) Size() int {
	return len(q.Elements)
}
