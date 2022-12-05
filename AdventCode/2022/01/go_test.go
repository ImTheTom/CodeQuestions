package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "first",
			args: "text.txt",
			want: 24000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputValue, _ := readFile(tt.args)

			assert.Equal(t, tt.want, DoQuestion(inputValue))
		})
	}
}

func Test_readFile(t *testing.T) {
	lines, err := readFile("test_file.txt")
	assert.NoError(t, err)
	assert.Equal(t, []string{
		"4601",
		"1583",
		"",
		"2995",
		"5319",
	}, lines)
}
