package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func DoQuestion(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	fmt.Println("Running main")
	node4 := &ListNode{
		Val: 9,
	}
	node3 := &ListNode{
		Val:  1,
		Next: node4,
	}
	node2 := &ListNode{
		Val:  5,
		Next: node3,
	}
	node := &ListNode{
		Val:  4,
		Next: node2,
	}
	tmpHead := node
	fmt.Println("Currently node list goes")
	for tmpHead != nil {
		fmt.Println(tmpHead.Val)
		tmpHead = tmpHead.Next
	}
	DoQuestion(node2)
	tmpHead = node
	fmt.Println("Now the node list goes")
	for tmpHead != nil {
		fmt.Println(tmpHead.Val)
		tmpHead = tmpHead.Next
	}
}
