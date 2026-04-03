package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func DoQuestion(head *ListNode) *ListNode {
	headPtr := head

	for {
		if head == nil || head.Next == nil {
			break
		}

		nextHead := head.Next

		total := head.Val + nextHead.Val

		highest := total / 2

		for {
			if highest == 0 {
				break
			}

			if head.Val%highest == 0 && nextHead.Val%highest == 0 {
				break
			}

			highest--
		}

		head.Next = &ListNode{
			Val:  highest,
			Next: nextHead,
		}

		head = nextHead
	}

	return headPtr
}

func main() {
	start := time.Now()

	l4 := &ListNode{
		Val: 3,
	}

	l3 := &ListNode{
		Val:  10,
		Next: l4,
	}

	l2 := &ListNode{
		Val:  6,
		Next: l3,
	}

	l1 := &ListNode{
		Val:  18,
		Next: l2,
	}

	fmt.Println("Running main")
	result := DoQuestion(l1)
	fmt.Printf("Got %v\n", result)

	i := 1

	for {
		if result == nil {
			break
		}

		fmt.Printf("Node %d - Val %d\n", i, result.Val)

		result = result.Next
	}

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
