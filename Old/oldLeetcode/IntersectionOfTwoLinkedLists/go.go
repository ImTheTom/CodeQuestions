package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func DoQuestion(headA, headB *ListNode) *ListNode {
	tempHeadAReset, tempHeadBeset := headA, headB
	aLen, bLen := getLengthOfList(headA), getLengthOfList(headB)
	headA, headB = tempHeadAReset, tempHeadBeset
	aDiff, bDiff := aLen-bLen, bLen-aLen

	for i := 0; i < aDiff; i++ {
		headA = headA.Next
	}
	for i := 0; i < bDiff; i++ {
		headB = headB.Next
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

func getLengthOfList(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func main() {
	fmt.Println("Running main")
	node5 := &ListNode{
		Val: 5,
	}
	node4 := &ListNode{
		Val:  4,
		Next: node5,
	}
	node8 := &ListNode{
		Val:  8,
		Next: node4,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node8,
	}
	node4Top := &ListNode{
		Val:  4,
		Next: node1,
	}
	node1Bot := &ListNode{
		Val:  1,
		Next: node8,
	}
	node6 := &ListNode{
		Val:  6,
		Next: node1Bot,
	}
	node5Bot := &ListNode{
		Val:  5,
		Next: node6,
	}
	result := DoQuestion(node4Top, node5Bot)
	for result != nil {
		fmt.Printf("Got %v\n", result.Val)
		result = result.Next
	}
}
