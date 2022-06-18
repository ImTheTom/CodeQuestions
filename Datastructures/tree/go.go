package main

import (
	"fmt"
	"time"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (b *BinarySearchTree) Insert(value int) {
	if b.Root == nil {
		b.Root = &Node{
			Value: value,
		}
		return
	}

	b.recursiveInsert(value, b.Root)
}

func (b *BinarySearchTree) recursiveInsert(value int, node *Node) {
	// Should never get here, but good to have.
	if node == nil {
		return
	}

	// Should value go to the left?
	if value <= node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value}

			return
		}

		b.recursiveInsert(value, node.Left)

		return
	}

	// Will continue here if going to the right
	if node.Right == nil {
		node.Right = &Node{Value: value}

		return
	}

	b.recursiveInsert(value, node.Right)
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.Value)
		printPreOrder(n.Left)
		printPreOrder(n.Right)
	}
}

func main() {
	start := time.Now()

	fmt.Println("Running main")

	t := &BinarySearchTree{}

	t.Insert(20)
	t.Insert(10)
	t.Insert(30)
	t.Insert(40)
	t.Insert(50)
	t.Insert(5)
	t.Insert(2)
	t.Insert(7)
	t.Insert(35)
	t.Insert(45)
	t.Insert(15)

	printPreOrder(t.Root)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
