package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first TreeNode
	}

	node7 := TreeNode{
		Val: 7,
	}
	node15 := TreeNode{
		Val: 15,
	}
	node9 := TreeNode{
		Val: 9,
	}
	node20 := TreeNode{
		Val:   20,
		Left:  &node15,
		Right: &node7,
	}
	node3 := TreeNode{
		Val:   3,
		Left:  &node9,
		Right: &node20,
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				first: node3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(&tt.args.first))
		})
	}
}
