package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{
				first: 3,
			},
			want: "III",
		},
		{
			name: "Second",
			args: args{
				first: 58,
			},
			want: "LVIII",
		},
		{
			name: "third",
			args: args{
				first: 1994,
			},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}
