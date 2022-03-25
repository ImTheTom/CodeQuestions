package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	len  int
}

func DoQuestion(head *ListNode) *ListNode {
	l := &LinkedList{}
	for head != nil {
		if head.Next == nil {
			l.Insert(head.Val)
			head = head.Next
			continue
		}
		l.Insert(head.Next.Val)
		l.Insert(head.Val)
		head = head.Next
		head = head.Next
	}
	return l.head
}

func (l *LinkedList) Insert(val int) {
	n := ListNode{}
	n.Val = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.Next == nil {
			ptr.Next = &n
			l.len++
			return
		}
		ptr = ptr.Next
	}
}

func main() {
	fmt.Println("Running main")

	l4 := &ListNode{
		Val:  4,
		Next: nil,
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

	result := DoQuestion(l1)
	for result != nil {
		fmt.Printf("Got %v\n", result)
		result = result.Next
	}
}
