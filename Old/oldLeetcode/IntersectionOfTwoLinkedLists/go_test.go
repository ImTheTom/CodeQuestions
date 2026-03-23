package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first  *ListNode
		second *ListNode
	}

	node5 := &ListNode{
		Val: 5,
	}
	node4 := &ListNode{
		Val:  4,
		Next: node5,
	}
	node8 := &ListNode{
		Val:  8,
		Next: node4,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node8,
	}
	node4Top := &ListNode{
		Val:  4,
		Next: node1,
	}
	node1Bot := &ListNode{
		Val:  1,
		Next: node8,
	}
	node6 := &ListNode{
		Val:  6,
		Next: node1Bot,
	}
	node5Bot := &ListNode{
		Val:  5,
		Next: node6,
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "first",
			args: args{
				first:  node4Top,
				second: node5Bot,
			},
			want: node8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first, tt.args.second))
		})
	}
}

func TestDoQuestionSecond(t *testing.T) {
	type args struct {
		first  *ListNode
		second *ListNode
	}

	node2 := &ListNode{
		Val: 2,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "first",
			args: args{
				first:  node1,
				second: node2,
			},
			want: node2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first, tt.args.second))
		})
	}
}
