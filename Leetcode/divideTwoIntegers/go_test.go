package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first  int
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
				first:  10,
				second: 3,
			},
			want: 3,
		},
		{
			name: "Second",
			args: args{
				first:  -2147483648,
				second: -1,
			},
			want: 2147483647,
		},
		{
			name: "Third",
			args: args{
				first:  -2147483648,
				second: 1,
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first, tt.args.second))
		})
	}
}
