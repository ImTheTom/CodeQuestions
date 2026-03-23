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
		want int
	}{
		{
			name: "first",
			args: args{
				first: []int{
					-2, 1, -3, 4, -1, 2, 1, -5, 4,
				},
			},
			want: 6,
		},
		{
			name: "second",
			args: args{
				first: []int{
					1,
				},
			},
			want: 1,
		},
		{
			name: "third",
			args: args{
				first: []int{
					5, 4, -1, 7, 8,
				},
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}
