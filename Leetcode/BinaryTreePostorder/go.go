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
	res = postorder(root, res)
	return res
}

func postorder(node *TreeNode, res []int) []int {
	if node != nil {
		res = postorder(node.Left, res)
		res = postorder(node.Right, res)
		res = append(res, node.Val)
	}
	return res
}

func main() {

	leftNode2 := &TreeNode{
		Val: 3,
	}

	rightNode1 := &TreeNode{
		Val:  2,
		Left: leftNode2,
	}

	root := &TreeNode{
		Val:   1,
		Right: rightNode1,
	}

	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion(root)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
