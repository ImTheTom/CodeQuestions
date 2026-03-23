package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				input: "III",
			},
			want: 3,
		},
		{
			name: "second",
			args: args{
				input: "LVIII",
			},
			want: 58,
		},
		{
			name: "third",
			args: args{
				input: "MCMXCIV",
			},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.input))
		})
	}
}
