package tree

import "Trees/node"

/**
       0
	  / \
	1    2
   /\   /\
  3 4  5 6
*/
func ExampleOne() *node.TreeNode {
	// Left side
	node3 := &node.TreeNode{
		Val: 3,
	}
	node4 := &node.TreeNode{
		Val: 4,
	}
	node1 := &node.TreeNode{
		Val:   1,
		Left:  node3,
		Right: node4,
	}

	// Right side
	node5 := &node.TreeNode{
		Val: 5,
	}
	node6 := &node.TreeNode{
		Val: 6,
	}
	node2 := &node.TreeNode{
		Val:   2,
		Left:  node5,
		Right: node6,
	}

	// Root
	return &node.TreeNode{
		Val:   1,
		Left:  node1,
		Right: node2,
	}
}
