package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	leftNode2 = &TreeNode{
		Val: 3,
	}

	rightNode1 = &TreeNode{
		Val:  2,
		Left: leftNode2,
	}

	root = &TreeNode{
		Val:   1,
		Right: rightNode1,
	}
)

func TestDoQuestion(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "first",
			root: root,
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(root))
		})
	}
}
