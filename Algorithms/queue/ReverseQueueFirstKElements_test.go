package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseQueueFirstKElements(t *testing.T) {
	input := Queue{}
	input.Enqueue(10)
	input.Enqueue(20)
	input.Enqueue(30)
	input.Enqueue(40)
	input.Enqueue(50)
	input.Enqueue(60)
	input.Enqueue(70)
	input.Enqueue(80)
	input.Enqueue(90)
	input.Enqueue(100)

	output := Queue{}
	output.Enqueue(50)
	output.Enqueue(40)
	output.Enqueue(30)
	output.Enqueue(20)
	output.Enqueue(10)
	output.Enqueue(60)
	output.Enqueue(70)
	output.Enqueue(80)
	output.Enqueue(90)
	output.Enqueue(100)

	type args struct {
		first  int
		second Queue
	}
	tests := []struct {
		name string
		args args
		want Queue
	}{
		{
			name: "first",
			args: args{
				first:  5,
				second: input,
			},
			want: output,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseQueueFirstKElements(tt.args.first, &tt.args.second)
			assert.Equal(t, tt.want.Elements, tt.args.second.Elements)
		})
	}
}
