package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var shiftDownElementArray = []int{
	3, 2, 2, 3,
}

var shiftDownElementWantArray = []int{
	2, 2, 3, 3,
}

func TestDoQuestion(t *testing.T) {
	type args struct {
		first  []int
		second int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				first: []int{
					3, 2, 2, 3,
				},
				second: 3,
			},
			want: 2,
		},
		{
			name: "second",
			args: args{
				first: []int{
					0, 1, 2, 2, 3, 0, 4, 2,
				},
				second: 2,
			},
			want: 5,
		},
		{
			name: "third",
			args: args{
				first: []int{
					2, 2, 3,
				},
				second: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first, tt.args.second))
		})
	}
}

func TestShiftDownElements(t *testing.T) {
	type args struct {
		first  []int
		second int
		third  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				first:  shiftDownElementArray,
				second: 0,
				third:  3,
			},
			want: shiftDownElementWantArray,
		},
		{
			name: "second",
			args: args{
				first:  shiftDownElementWantArray,
				second: 2,
				third:  3,
			},
			want: shiftDownElementWantArray,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, shiftDownElements(tt.args.first, tt.args.second, tt.args.third))
		})
	}
}
