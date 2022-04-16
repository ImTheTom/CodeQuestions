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
				first: []int{
					0, 1, 0, 3, 12,
				},
			},
			want: []int{
				1, 3, 12, 0, 0,
			},
		},
		{
			name: "second",
			args: args{
				first: []int{
					0, 0, 1,
				},
			},
			want: []int{
				1, 0, 0,
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

func TestShiftDown(t *testing.T) {
	type args struct {
		first  int
		second int
		third  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				first:  0,
				second: 4,
				third: []int{
					0, 1, 0, 3, 12,
				},
			},
			want: []int{
				1, 0, 3, 12, 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shiftDown(tt.args.first, tt.args.second, tt.args.third)
			assert.Equal(t, tt.want, tt.args.third)
		})
	}
}
