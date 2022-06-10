package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first  [][]byte
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
				first: [][]byte{
					{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
				},
				second: "ABCCED",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first, tt.args.second))
		})
	}
}
