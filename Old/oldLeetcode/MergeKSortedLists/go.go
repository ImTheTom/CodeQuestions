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

func DoQuestion(lists []*ListNode) *ListNode {
	result := LinkedList{}

	moreToGo := true
	for moreToGo {
		smallest := 9999
		smallestIndex := -1

		for i, v := range lists {
			if v == nil {
				continue
			}

			if v.Val < smallest {
				smallest = v.Val
				smallestIndex = i
			}
		}

		if smallestIndex == -1 {
			moreToGo = false
		} else {
			result.Insert(lists[smallestIndex].Val)
			lists[smallestIndex] = lists[smallestIndex].Next
		}
	}

	return result.head
}

func main() {
	fmt.Println("Running main")

	l1 := LinkedList{}
	l2 := LinkedList{}
	l3 := LinkedList{}

	l1.Insert(1)
	l1.Insert(4)
	l1.Insert(5)

	l2.Insert(1)
	l2.Insert(3)
	l2.Insert(4)

	l3.Insert(2)
	l3.Insert(6)

	result := DoQuestion([]*ListNode{
		l1.head, l2.head, l3.head,
	})
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
