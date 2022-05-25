package tree

import (
	"Trees/node"
	"fmt"
)

type Tree struct {
	Root *node.TreeNode
}

func (t *Tree) string() {
	fmt.Println(t.Root.Val)
}
