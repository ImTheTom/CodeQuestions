package main

import (
	"fmt"
	"reflect"
	"testing"
)

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 74382},
	{input: 382399},
}

func BenchmarkNotLengthSpecified(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("Input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NotLengthSpecified(v.input)
			}
		})
	}
}

func BenchmarkLengthSpecified(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("Input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LengthSpecified(v.input)
			}
		})
	}
}

func TestNotLengthSpecified(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				size: 7438200,
			},
			want: 7438200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotLengthSpecified(tt.args.size); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("NotLengthSpecified() = %v, want %v", got, tt.want)
			}
		})
	}
}
