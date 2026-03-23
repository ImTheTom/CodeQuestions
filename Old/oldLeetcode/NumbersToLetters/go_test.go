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
		want []string
	}{
		{
			name: "first",
			args: args{
				first: "23",
			},
			want: []string{
				"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
			},
		},
		{
			name: "second",
			args: args{
				first: "",
			},
			want: []string{},
		},
		{
			name: "third",
			args: args{
				first: "2",
			},
			want: []string{
				"a", "b", "c",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}
