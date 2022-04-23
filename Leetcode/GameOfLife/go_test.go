package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "first",
			args: args{
				first: [][]int{
					{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoQuestion(tt.args.first)
			assert.Equal(t, tt.want, tt.args.first)
		})
	}
}
