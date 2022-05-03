package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				first: "leetcode",
			},
			want: 0,
		},
		{
			name: "second",
			args: args{
				first: "loveleetcode",
			},
			want: 2,
		},
		{
			name: "third",
			args: args{
				first: "aabb",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}
