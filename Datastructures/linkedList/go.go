package main

import (
	"fmt"
	"time"
)

// Three main types. Single, double and circular. Circular can be either double or single.

func main() {
	start := time.Now()

	fmt.Println("Running main")

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

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
