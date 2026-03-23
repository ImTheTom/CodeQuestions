package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first *ListNode
	}

	firstNodeNeg4 := &ListNode{
		Val: -4,
	}
	firstNode0 := &ListNode{
		Val:  0,
		Next: firstNodeNeg4,
	}
	firstNode2 := &ListNode{
		Val:  2,
		Next: firstNode0,
	}
	firstNodeNeg4.Next = firstNode2
	firstNode3 := &ListNode{
		Val:  3,
		Next: firstNode2,
	}

	secondNode2 := &ListNode{
		Val: 2,
	}
	secondNode1 := &ListNode{
		Val:  1,
		Next: secondNode2,
	}
	secondNode2.Next = secondNode1

	thirdNode2 := &ListNode{
		Val: 1,
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{
				first: firstNode3,
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				first: secondNode1,
			},
			want: true,
		},
		{
			name: "first",
			args: args{
				first: thirdNode2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}
