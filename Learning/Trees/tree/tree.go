package tree

import (
	"Trees/node"
	"fmt"
)

type BinaryTree struct {
	Root *node.BinaryNode
}

func (t *BinaryTree) string() {
	fmt.Println(t.Root.Val)
}
