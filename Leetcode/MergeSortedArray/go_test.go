package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShiftDownArray(t *testing.T) {
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
				first: []int{
					1, 2, 3, 4, 5, 6,
				},
				second: 1,
				third:  1,
			},
			want: []int{
				1, 1, 2, 3, 4, 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShiftDownArray(tt.args.first, tt.args.second, tt.args.third)
			assert.Equal(t, tt.want, tt.args.first)
		})
	}
}
