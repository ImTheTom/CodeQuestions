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

func DoQuestion(list1 *ListNode, list2 *ListNode) *ListNode {
	var full []int
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			full = append(full, list1.Val)
			list1 = list1.Next
		} else {
			full = append(full, list2.Val)
			list2 = list2.Next
		}
	}
	for list1 != nil {
		full = append(full, list1.Val)
		list1 = list1.Next
	}
	for list2 != nil {
		full = append(full, list2.Val)
		list2 = list2.Next
	}

	linkedList := LinkedList{}

	for _, v := range full {
		linkedList.Insert(v)
	}

	return linkedList.head
}

func main() {
	fmt.Println("Running main")

	l13 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l12 := &ListNode{
		Val:  2,
		Next: l13,
	}
	l11 := &ListNode{
		Val:  1,
		Next: l12,
	}

	l23 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l22 := &ListNode{
		Val:  3,
		Next: l23,
	}
	l21 := &ListNode{
		Val:  1,
		Next: l22,
	}

	result := DoQuestion(l11, l21)
	for result != nil {
		fmt.Printf("Got %v\n", result)
		result = result.Next
	}
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
