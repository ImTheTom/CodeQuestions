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
		want int
	}{
		{
			name: "first",
			args: args{
				first: []string{
					"abcw", "baz", "foo", "bar", "xtfn", "abcdef",
				},
			},
			want: 16,
		},
		{
			name: "second",
			args: args{
				first: []string{
					"a", "ab", "abc", "d", "cd", "bcd", "abcd",
				},
			},
			want: 4,
		},
		{
			name: "third",
			args: args{
				first: []string{
					"a", "aa", "aaa", "aaaa",
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoQuestion(tt.args.first))
		})
	}
}

func TestDoesNotShareSameCharacter(t *testing.T) {
	type args struct {
		first  string
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
				first:  "abc",
				second: "baz",
			},
			want: false,
		},
		{
			name: "second",
			args: args{
				first:  "abc",
				second: "tdz",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DoesNotShareSameCharacter(tt.args.first, tt.args.second))
		})
	}
}

func TestMaxValue(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				values: []int{
					1, 2, 3, 5, 4,
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MaxValue(tt.args.values))
		})
	}
}
