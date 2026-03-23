package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first  string
		second int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{
				first:  "abcd",
				second: 2,
			},
			want: "abcd",
		},
		{
			name: "second",
			args: args{
				first:  "deeedbbcccbdaa",
				second: 3,
			},
			want: "aa",
		},
		{
			name: "third",
			args: args{
				first:  "pbbcggttciiippooaais",
				second: 2,
			},
			want: "ps",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first, tt.args.second))
		})
	}
}
