package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := make([]*ListNode, 0)

	loopNode := head

	for {
		if loopNode == nil {
			break
		}

		nodes = append(nodes, loopNode)

		loopNode = loopNode.Next
	}

	nodesLength := len(nodes)
	nodeRemovalIdx := nodesLength - n

	var (
		nodeAhead  *ListNode
		nodeBehind *ListNode
	)

	if nodeRemovalIdx+1 < nodesLength {
		nodeAhead = nodes[nodeRemovalIdx+1]
	}

	if nodeRemovalIdx-1 >= 0 {
		nodeBehind = nodes[nodeRemovalIdx-1]
	}

	if nodeBehind != nil {
		nodeBehind.Next = nodeAhead
	} else {
		return nodes[nodeRemovalIdx].Next
	}

	return head
}

func main() {
	testCaseID := 2

	start := time.Now()

	l5 := &ListNode{
		Val: 5,
	}

	l4 := &ListNode{
		Val:  4,
		Next: l5,
	}

	l3 := &ListNode{
		Val:  3,
		Next: l4,
	}

	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}

	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}

	l21 := &ListNode{
		Val: 1,
	}

	l32 := &ListNode{
		Val: 2,
	}

	l31 := &ListNode{
		Val:  1,
		Next: l32,
	}

	testCase := l1
	switch testCaseID {
	case 2:
		testCase = l21
	case 3:
		testCase = l31
	}

	fmt.Println("Running main")
	result := removeNthFromEnd(testCase, 1)

	i := 0

	for {
		if result == nil {
			break
		}

		fmt.Printf("Node %d - Val %d\n", i, result.Val)

		i++
		result = result.Next
	}

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
