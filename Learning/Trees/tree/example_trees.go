package tree

import "Trees/node"

/**
       0
	  / \
	1    2
   /\   /\
  3 4  5 6
*/
func BinaryExampleOne() *node.BinaryNode {
	// Left side
	node3 := &node.BinaryNode{
		Val: 3,
	}
	node4 := &node.BinaryNode{
		Val: 4,
	}
	node1 := &node.BinaryNode{
		Val:   1,
		Left:  node3,
		Right: node4,
	}

	// Right side
	node5 := &node.BinaryNode{
		Val: 5,
	}
	node6 := &node.BinaryNode{
		Val: 6,
	}
	node2 := &node.BinaryNode{
		Val:   2,
		Left:  node5,
		Right: node6,
	}

	// Root
	return &node.BinaryNode{
		Val:   1,
		Left:  node1,
		Right: node2,
	}
}
