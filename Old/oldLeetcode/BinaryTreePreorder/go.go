package main

import (
	"fmt"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DoQuestion(root *TreeNode) []int {
	res := make([]int, 0)
	res = preOrder(root, res)
	return res
}

func preOrder(root *TreeNode, res []int) []int {
	if root != nil {
		res = append(res, root.Val)
		res = preOrder(root.Left, res)
		res = preOrder(root.Right, res)
	}
	return res
}

func main() {
	start := time.Now()

	leftNode1 := &TreeNode{
		Val: 1,
	}

	rightNode1 := &TreeNode{
		Val: 3,
	}

	root := &TreeNode{
		Val:   2,
		Left:  leftNode1,
		Right: rightNode1,
	}

	fmt.Println("Running main")
	result := DoQuestion(root)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
