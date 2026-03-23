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
					5, 7, 7, 8, 8, 10,
				},
				second: 8,
			},
			want: []int{
				3, 4,
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
				0, 0,
			},
		},
		{
			name: "third",
			args: args{
				first: []int{
					1, 1, 2,
				},
				second: 1,
			},
			want: []int{
				0, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first, tt.args.second))
		})
	}
}
