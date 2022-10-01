package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want []int
	}{
		{
			name: "first",
			root: root,
			want: []int{5, 6, 3, 2, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.root))
		})
	}
}
