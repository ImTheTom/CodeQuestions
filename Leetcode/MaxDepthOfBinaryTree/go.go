package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DoQuestion(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := DoQuestion(root.Left)
	rightDepth := DoQuestion(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main() {
	fmt.Println("Running main")
	node7 := TreeNode{
		Val: 7,
	}
	node15 := TreeNode{
		Val: 15,
	}
	node9 := TreeNode{
		Val: 9,
	}
	node20 := TreeNode{
		Val:   20,
		Left:  &node15,
		Right: &node7,
	}
	node3 := TreeNode{
		Val:   3,
		Left:  &node9,
		Right: &node20,
	}
	result := DoQuestion(&node3)
	fmt.Printf("Got %v\n", result)
}
