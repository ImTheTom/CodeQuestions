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
				first: []int{
					1, 1, 1, 2, 2, 3,
				},
				second: 2,
			},
			want: []int{
				1, 2,
			},
		},
		{
			name: "second",
			args: args{
				first: []int{
					1,
				},
				second: 1,
			},
			want: []int{
				1,
			},
		},
		{
			name: "third",
			args: args{
				first: []int{
					3, 0, 1, 0,
				},
				second: 1,
			},
			want: []int{
				0,
			},
		},
		{
			name: "fourth",
			args: args{
				first: []int{
					5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3,
				},
				second: 3,
			},
			want: []int{
				3, 7, 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first, tt.args.second))
		})
	}
}
