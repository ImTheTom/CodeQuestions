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
				first:  "PAYPALISHIRING",
				second: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first, tt.args.second))
		})
	}
}
