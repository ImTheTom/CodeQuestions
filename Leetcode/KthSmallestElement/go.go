package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DoQuestion(root *TreeNode, k int) int {
	values := []int{}

	exploreTree(root, &values)

	sort.Ints(values)

	return values[k-1]
}

func exploreTree(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}

	*values = append(*values, root.Val)

	exploreTree(root.Left, values)
	exploreTree(root.Right, values)
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion(buildTree(), 1)
	fmt.Printf("Got %v\n", result)
}

func buildTree() *TreeNode {
	// Left side
	two := &TreeNode{
		Val: 2,
	}
	one := &TreeNode{
		Val:   1,
		Right: two,
	}

	// Right side
	four := &TreeNode{
		Val: 4,
	}

	// Root
	return &TreeNode{
		Val:   3,
		Left:  one,
		Right: four,
	}
}
