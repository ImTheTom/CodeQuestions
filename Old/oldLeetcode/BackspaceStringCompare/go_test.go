package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoQuestion(t *testing.T) {
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
				first:  "ab#c",
				second: "ad#c",
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				first:  "ab##",
				second: "c#d#",
			},
			want: true,
		},
		{
			name: "third",
			args: args{
				first:  "a#c",
				second: "b",
			},
			want: false,
		},
		{
			name: "fourth",
			args: args{
				first:  "xj##tw",
				second: "bxo#j##tw",
			},
			want: false,
		},
		{
			name: "five",
			args: args{
				first:  "a##c",
				second: "#a#c",
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

func Test_stripHashtags(t *testing.T) {
	type args struct {
		first string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first",
			args: args{
				first: "ab#c",
			},
			want: "ac",
		},
		{
			name: "second",
			args: args{
				first: "ab##",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, stripHashtags(tt.args.first))
		})
	}
}
