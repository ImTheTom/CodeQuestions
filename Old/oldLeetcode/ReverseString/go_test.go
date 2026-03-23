package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "first",
			args: args{
				first: []byte{
					'h', 'e', 'l', 'l', 'o',
				},
			},
			want: []byte{
				'o', 'l', 'l', 'e', 'h',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoQuestion(tt.args.first)
			assert.Equal(t, tt.want, tt.args.first)
		})
	}
}
