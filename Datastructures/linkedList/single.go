package main

import (
	"errors"
	"fmt"
)

type SingleNode struct {
	Val  int
	Next *SingleNode
}

type SingleLinkedList struct {
	Head *SingleNode
}

func (l *SingleLinkedList) Search(value int) *SingleNode {
	tmpNode := l.Head

	for tmpNode != nil {
		if tmpNode.Val == value {
			return tmpNode
		}

		tmpNode = tmpNode.Next
	}

	return nil
}

func (l *SingleLinkedList) Append(value int) {
	newNode := &SingleNode{
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
}

func (l *SingleLinkedList) InsertToHead(value int) {
	newNode := &SingleNode{
		Val:  value,
		Next: l.Head,
	}

	l.Head = newNode
}

func (l *SingleLinkedList) InsertAtIndex(value, index int) error {
	if index == 0 {
		l.InsertToHead(value)
		return nil
	}

	newNode := &SingleNode{
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

	return nil
}

func (l *SingleLinkedList) RemoveHead() {
	if l.Head == nil {
		return
	}

	tmpNewHead := l.Head.Next

	l.Head = nil

	l.Head = tmpNewHead
}

func (l *SingleLinkedList) RemoveEnd() {
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

func (l *SingleLinkedList) RemoveAtIndex(index int) error {
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
	tmpNode = nil

	return nil
}

func (l *SingleLinkedList) String() {
	var (
		index int
		final string
	)

	tmpNode := l.Head

	for tmpNode != nil {
		final = fmt.Sprintf("%sIndex - %d Value - %d\n", final, index, tmpNode.Val)
		index++
		tmpNode = tmpNode.Next
	}

	fmt.Println(final)
}
