package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	type args struct {
		nums          []int
		rotationIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "First case",
			args: args{
				nums: []int{
					1, 2, 3, 4, 5, 6, 7,
				},
				rotationIndex: 2,
			},
			want: []int{
				3, 4, 5, 6, 7, 1, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Rotate(tt.args.nums, tt.args.rotationIndex))
		})
	}
}
