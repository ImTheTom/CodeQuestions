package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
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

	type args struct {
		first *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				first: nil,
			},
			want: []int{},
		},
		{
			name: "second",
			args: args{
				first: root,
			},
			want: []int{
				1, 3, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}
