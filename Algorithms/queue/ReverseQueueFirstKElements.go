package main

func ReverseQueueFirstKElements(k int, q *Queue) {
	if q.Empty() || k > q.Size() || k <= 0 {
		return
	}

	s := Stack{}

	for i := 0; i < k; i++ {
		s.Push(q.Dequeue())
	}

	for !s.IsEmpty() {
		q.Enqueue(s.Pop())
	}

	for i := 0; i < q.Size()-k; i++ {
		q.Enqueue(q.Dequeue())
	}
}
