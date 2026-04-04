package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	resultHead := &ListNode{}

	actualStart := head.Next

	resultWork := resultHead

	currentTotal := 0

	for {
		if actualStart == nil {
			break
		}

		if actualStart.Val != 0 {
			currentTotal += actualStart.Val

			actualStart = actualStart.Next

			continue
		}

		if resultWork.Val == 0 {
			initialHead := &ListNode{
				Val: currentTotal,
			}
			resultHead = initialHead

			resultWork = initialHead

			currentTotal = 0

			actualStart = actualStart.Next

			continue
		}

		newNode := &ListNode{
			Val: currentTotal,
		}

		resultWork.Next = newNode

		resultWork = newNode

		actualStart = actualStart.Next

		currentTotal = 0
	}

	return resultHead
}

func main() {
	start := time.Now()

	l8 := &ListNode{
		Val: 0,
	}

	l7 := &ListNode{
		Val:  2,
		Next: l8,
	}

	l6 := &ListNode{
		Val:  5,
		Next: l7,
	}

	l5 := &ListNode{
		Val:  4,
		Next: l6,
	}

	l4 := &ListNode{
		Val:  0,
		Next: l5,
	}

	l3 := &ListNode{
		Val:  1,
		Next: l4,
	}

	l2 := &ListNode{
		Val:  3,
		Next: l3,
	}

	l1 := &ListNode{
		Val:  0,
		Next: l2,
	}

	fmt.Println("Running main")
	result := mergeNodes(l1)
	fmt.Printf("Got %v\n", result)

	nodeIdx := 1

	for {
		if result == nil {
			break
		}

		fmt.Printf("Node - %d Val - %d\n", nodeIdx, result.Val)

		nodeIdx++
		result = result.Next
	}

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
