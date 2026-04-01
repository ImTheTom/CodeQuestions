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
		want bool
	}{
		{
			name: "first",
			args: args{
				first: "()[]{}",
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				first: "([])",
			},
			want: true,
		},
		{
			name: "third",
			args: args{
				first: "([)]",
			},
			want: false,
		},
		{
			name: "fourth",
			args: args{
				first: "(",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}
