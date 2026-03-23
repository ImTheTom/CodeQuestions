package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
	End  *ListNode
}

func DoQuestion(head *ListNode) *ListNode {
	l := &List{}
	for head != nil {
		tmpHead := head
		var previous *ListNode = nil
		for tmpHead.Next != nil {
			previous = tmpHead
			tmpHead = tmpHead.Next
		}
		if previous != nil {
			previous.Next = nil
		} else {
			head = nil
		}
		l.Insert(tmpHead.Val)
	}
	return l.Head
}

func (l *List) Insert(val int) {
	node := &ListNode{
		Val: val,
	}
	if l.Head == nil {
		l.Head = node
		l.End = node
	} else {
		l.End.Next = node
		l.End = node
	}
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
	node3 := &ListNode{
		Val:  3,
		Next: node4,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}
	result := DoQuestion(node1)
	for result != nil {
		fmt.Printf("Got %v\n", result.Val)
		result = result.Next
	}
}
