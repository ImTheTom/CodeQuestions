package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

const maxNodes = 10000

func DoQuestion(head *ListNode) bool {
	total := 0

	for head != nil {
		head = head.Next
		total++

		if total > maxNodes {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("Running main")
	nodeNeg4 := &ListNode{
		Val: -4,
	}
	node0 := &ListNode{
		Val:  0,
		Next: nodeNeg4,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node0,
	}
	nodeNeg4.Next = node2
	node3 := &ListNode{
		Val:  3,
		Next: node2,
	}
	result := DoQuestion(node3)
	fmt.Printf("Got %v\n", result)
}
