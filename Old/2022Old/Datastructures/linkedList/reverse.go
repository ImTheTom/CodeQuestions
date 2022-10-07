package linkedlist

import (
	"errors"
	"fmt"
)

type ReverseNode struct {
	Val  int
	Next *ReverseNode
	Prev *ReverseNode
}

type ReverseLinkedList struct {
	Head *ReverseNode
}

func (l *ReverseLinkedList) Search(value int) *ReverseNode {
	tmpNode := l.Head

	for tmpNode != nil {
		if tmpNode.Val == value {
			return tmpNode
		}

		tmpNode = tmpNode.Next
	}

	return nil
}

func (l *ReverseLinkedList) Append(value int) {
	newNode := &ReverseNode{
		Val: value,
	}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	tmpNode := l.Head
	for tmpNode.Next != nil {
		tmpNode = tmpNode.Next
	}

	tmpNode.Next = newNode
	newNode.Prev = tmpNode
}

func (l *ReverseLinkedList) InsertToHead(value int) {
	newNode := &ReverseNode{
		Val:  value,
		Next: l.Head,
		Prev: nil,
	}

	l.Head.Prev = newNode
	l.Head = newNode
}

func (l *ReverseLinkedList) InsertAtIndex(value, index int) error {
	if index == 0 {
		l.InsertToHead(value)
		return nil
	}

	newNode := &ReverseNode{
		Val: value,
	}

	tmpNode := l.Head
	for i := 0; i < index-1; i++ {
		if tmpNode == nil {
			return errors.New("Attempting to insert at Index greater than length")
		}
		tmpNode = tmpNode.Next
	}

	tmpNodeNext := tmpNode.Next

	tmpNode.Next = newNode

	newNode.Next = tmpNodeNext
	newNode.Prev = tmpNode

	tmpNodeNext.Prev = newNode

	return nil
}

func (l *ReverseLinkedList) RemoveHead() {
	if l.Head == nil {
		return
	}

	tmpNewHead := l.Head.Next

	l.Head = nil

	l.Head = tmpNewHead
	l.Head.Prev = nil
}

func (l *ReverseLinkedList) RemoveEnd() {
	if l.Head == nil {
		return
	}

	tmpNode := l.Head
	previousNode := l.Head

	for tmpNode.Next != nil {
		previousNode = tmpNode
		tmpNode = tmpNode.Next
	}

	tmpNode = nil
	previousNode.Next = nil
}

func (l *ReverseLinkedList) RemoveAtIndex(index int) error {
	if index == 0 {
		l.RemoveHead()
		return nil
	}

	tmpNode := l.Head
	previousNode := l.Head

	for i := 0; i < index-1; i++ {
		if tmpNode == nil {
			return errors.New("Attempting to insert at Index greater than length")
		}
		previousNode = tmpNode
		tmpNode = tmpNode.Next
	}

	previousNode.Next = tmpNode.Next
	tmpNode.Next.Prev = tmpNode.Prev
	tmpNode = nil

	return nil
}

func (l *ReverseLinkedList) String() {
	var (
		index int
		final string
	)

	tmpNode := l.Head

	for tmpNode != nil {
		if tmpNode.Prev != nil {
			final = fmt.Sprintf("%sIndex - %d Value - %d Previous Value - %d\n", final, index, tmpNode.Val, tmpNode.Prev.Val)
		} else {
			final = fmt.Sprintf("%sIndex - %d Value - %d\n", final, index, tmpNode.Val)
		}
		index++
		tmpNode = tmpNode.Next
	}

	fmt.Println(final)
}
