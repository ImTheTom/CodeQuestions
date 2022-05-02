package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isArraySortedByParity(arr []int) bool {
	doneWithEvens := false
	for _, v := range arr {
		if !doneWithEvens {
			if v%2 != 0 {
				doneWithEvens = true
			}
		} else {
			if v%2 == 0 {
				return false
			}
		}
	}
	return true
}

func TestDoQuestion(t *testing.T) {
	type args struct {
		first []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "first",
			args: args{
				first: []int{
					3, 1, 2, 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := DoQuestion(tt.args.first)
			if len(res) != len(tt.args.first) {
				t.Log("Different lengths")
				t.Fail()
			}
			assert.True(t, isArraySortedByParity(res), "Array is not parity sorted %v", res)
			sort.Ints(res)
			sort.Ints(tt.args.first)
			assert.Equal(t, tt.args.first, res, "Elements don't match")
		})
	}
}
