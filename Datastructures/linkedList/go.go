package main

import (
	"fmt"
	"time"
)

// Three main types. Single, double and circular. Circular can be either double or single.

func main() {
	start := time.Now()

	fmt.Println("Running main")

	DoReverseList()

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}

func DoSingleList() {
	l := &SingleLinkedList{}

	l.Append(5)
	l.Append(6)
	l.Append(7)

	l.String()

	l.InsertToHead(2)

	l.InsertAtIndex(3, 1)

	l.String()

	l.RemoveHead()
	l.RemoveEnd()
	l.RemoveAtIndex(2)

	l.String()

	l.Append(8)
	l.Append(9)
	l.Append(10)

	node := l.Search(9)
	fmt.Println(node)

	node = l.Search(39)
	fmt.Println(node)
}

func DoReverseList() {
	l := &ReverseLinkedList{}

	l.Append(5)
	l.Append(6)
	l.Append(7)

	l.String()

	l.InsertToHead(2)

	l.InsertAtIndex(3, 1)

	l.String()

	l.RemoveHead()
	l.RemoveEnd()
	l.RemoveAtIndex(2)

	l.String()

	l.Append(8)
	l.Append(9)
	l.Append(10)

	node := l.Search(9)
	fmt.Println(node)

	node = l.Search(39)
	fmt.Println(node)
}
