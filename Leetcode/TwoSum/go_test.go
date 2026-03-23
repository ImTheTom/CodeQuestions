package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first  []int
		second int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				first:  []int{2, 7, 11, 15},
				second: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "second",
			args: args{
				first:  []int{3, 2, 4},
				second: 6,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion2(tt.args.first, tt.args.second))
		})
	}
}
