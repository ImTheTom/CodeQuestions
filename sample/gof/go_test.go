package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
	type args struct {
		first  int
		second int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				first:  1,
				second: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first, tt.args.second))
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
