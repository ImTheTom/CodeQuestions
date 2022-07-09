package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreBracketsBalanced(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       bool
	}{
		{
			name:       "balanced",
			expression: "{()}[]",
			want:       true,
		},
		{
			name:       "unbalanced",
			expression: "{()}[",
			want:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, AreBracketsBalanced(tt.expression))
		})
	}
}
