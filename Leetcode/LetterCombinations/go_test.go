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
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}
