package main

type ReverseNode struct {
	Val  int
	Next *ReverseNode
	Prev *ReverseNode
}

type ReverseLinkedList struct {
	Head *ReverseNode
}
