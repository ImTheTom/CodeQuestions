package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	tests := []struct {
		name  string
		edges [][]int
		want  int
	}{
		{
			name: "first",
			edges: [][]int{
				{1, 2}, {2, 3}, {4, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.edges))
		})
	}
}
