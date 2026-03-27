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
				first: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "first",
			args: args{
				first: "bbbbb",
			},
			want: 1,
		},
		{
			name: "first",
			args: args{
				first: "pwwkew",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}
