package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				first: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "second",
			args: args{
				first: []int{-1, 1, 0, -3, 3},
			},
			want: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}
