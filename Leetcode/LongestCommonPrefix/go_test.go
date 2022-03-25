package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{
				first: []string{
					"flower", "flow", "flight",
				},
			},
			want: "fl",
		},
		{
			name: "second",
			args: args{
				first: []string{
					"dog", "racecar", "car",
				},
			},
			want: "",
		},
		{
			name: "third",
			args: args{
				first: []string{
					"ab", "a",
				},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}
