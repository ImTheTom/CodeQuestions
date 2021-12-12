package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	len  int
}

var l13 = ListNode{
	Val:  3,
	Next: nil,
}

var l12 = ListNode{
	Val:  4,
	Next: &l13,
}

var l11 = ListNode{
	Val:  2,
	Next: &l12,
}

var l23 = ListNode{
	Val:  4,
	Next: nil,
}

var l22 = ListNode{
	Val:  6,
	Next: &l23,
}

var l21 = ListNode{
	Val:  5,
	Next: &l22,
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ll := LinkedList{
		head: nil,
		len:  0,
	}

	previous := 0
	moreL1 := true
	moreL2 := true
	total := 0

	for {
		total, moreL1, moreL2 = sumValuesAndProgress(l1, l2, moreL1, moreL2)
		total += previous
		if total > 9 {
			total -= 10
			previous = 1
		} else {
			previous = 0
		}

		ll.insert(total)

		total = 0

		if !moreL1 && !moreL2 {
			break
		}
	}

	if previous == 1 {
		ll.insert(1)
	}

	return ll.head
}

func sumValuesAndProgress(l1 *ListNode, l2 *ListNode, moreL1, moreL2 bool) (int, bool, bool) {
	total := 0

	if moreL1 {
		total += l1.Val
		if l1.Next != nil {
			*l1 = *l1.Next
		} else {
			moreL1 = false
		}
	}

	if moreL2 {
		total += l2.Val
		if l2.Next != nil {
			*l2 = *l2.Next
		} else {
			moreL2 = false
		}
	}

	return total, moreL1, moreL2
}

func (l *LinkedList) insert(val int) {
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

	resultNode := addTwoNumbers(&l11, &l21)

	if resultNode == nil {
		fmt.Println("Result Node was nil")
	} else {
		var s string
		for {
			s = fmt.Sprintf("%s%d", s, resultNode.Val)
			if resultNode.Next == nil {
				break
			}
			resultNode = resultNode.Next
		}
		fmt.Printf("%s\n", s)
	}
}
