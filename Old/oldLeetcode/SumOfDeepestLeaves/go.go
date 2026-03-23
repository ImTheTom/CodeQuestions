package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DoQuestion(root *TreeNode) int {
	result := make(map[int]int)

	loadValues(root, 0, result)

	max, maxIndex := -1, -1

	for key, val := range result {
		if key > maxIndex {
			maxIndex = key
			max = val
		}
	}

	return max
}

func loadValues(node *TreeNode, key int, res map[int]int) {
	if node == nil {
		return
	}

	curr := res[key]
	res[key] = curr + node.Val

	loadValues(node.Left, key+1, res)
	loadValues(node.Right, key+1, res)
}

func buildTree() *TreeNode {
	// Left side
	seven := &TreeNode{
		Val: 7,
	}
	four := &TreeNode{
		Val:  4,
		Left: seven,
	}
	five := &TreeNode{
		Val: 5,
	}
	two := &TreeNode{
		Val:   2,
		Left:  four,
		Right: five,
	}

	// Right side
	eight := &TreeNode{
		Val: 8,
	}
	six := &TreeNode{
		Val:   6,
		Right: eight,
	}
	three := &TreeNode{
		Val:   3,
		Right: six,
	}

	// Root node
	return &TreeNode{
		Val:   1,
		Left:  two,
		Right: three,
	}
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion(buildTree())
	fmt.Printf("Got %v\n", result)
}
