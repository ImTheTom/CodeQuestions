package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	leftNode1 = &TreeNode{
		Val: 1,
	}

	rightNode1 = &TreeNode{
		Val: 3,
	}

	root = &TreeNode{
		Val:   2,
		Left:  leftNode1,
		Right: rightNode1,
	}
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "general",
			args: args{
				root: root,
			},
			want: []int{
				2, 1, 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.root))
		})
	}
}

func BenchmarkDoQuestion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DoQuestion(root)
	}
}
