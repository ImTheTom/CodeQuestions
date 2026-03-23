package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{
				first:  "anagram",
				second: "nagaram",
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				first:  "cat",
				second: "car",
			},
			want: false,
		},
		{
			name: "third",
			args: args{
				first:  "a",
				second: "ab",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first, tt.args.second))
		})
	}
}
