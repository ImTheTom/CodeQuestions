package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want string
	}{
		{
			name: "first",
			ip:   "1.1.1.1",
			want: "1[.]1[.]1[.]1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.ip))
		})
	}
}
