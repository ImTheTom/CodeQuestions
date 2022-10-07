package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	type args struct {
		arr []int
		pos int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				arr: []int{
					1, 2, 3, 4, 5, 6, 7,
				},
				pos: 2,
			},
			want: []int{
				3, 4, 5, 6, 7, 1, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate(tt.args.arr, tt.args.pos)
			assert.Equal(t, tt.want, tt.args.arr)
		})
	}
}
