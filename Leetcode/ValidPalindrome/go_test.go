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
				first: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				first: "race a car",
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
