package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "first",
			args: []int{
				1, 2, 1,
			},
			want: []int{
				1, 2, 1, 1, 2, 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args))
		})
	}
}
