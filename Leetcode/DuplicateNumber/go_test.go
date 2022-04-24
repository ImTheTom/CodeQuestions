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
					1, 3, 4, 2, 2,
				},
			},
			want: 2,
		},
		{
			name: "second",
			args: args{
				first: []int{
					3, 1, 3, 4, 2,
				},
			},
			want: 3,
		},
		{
			name: "third",
			args: args{
				first: []int{
					2, 2, 2, 2, 2,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}
