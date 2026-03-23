package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func DoQuestion(root *TreeNode) []int {
	result := &[]int{}
	inorder(root, result)
	return *result
}

func inorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)
}

func main() {
	fmt.Println("Running main")
	branch3 := &TreeNode{
		Val: 3,
	}
	branch2 := &TreeNode{
		Val:   2,
		Left:  branch3,
		Right: nil,
	}
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: branch2,
	}
	result := DoQuestion(root)
	fmt.Printf("Got %v\n", result)
}
