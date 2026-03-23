package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	type args struct {
		words  []string
		search string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{
				words:  []string{"NO", "YES"},
				search: "YES",
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				words:  []string{"NO", "YES"},
				search: "MAYBE",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := Constructor()
			for _, w := range tt.args.words {
				trie.Insert(w)
			}
			assert.Equal(t, tt.want, trie.Search(tt.args.search))
		})
	}
}

func TestStartsWith(t *testing.T) {
	type args struct {
		words  []string
		prefix string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "first",
			args: args{
				words:  []string{"Something"},
				prefix: "Some",
			},
			want: true,
		},
		{
			name: "second",
			args: args{
				words:  []string{"Something"},
				prefix: "No",
			},
			want: false,
		},
		{
			name: "third",
			args: args{
				words:  []string{"some"},
				prefix: "something",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := Constructor()
			for _, w := range tt.args.words {
				trie.Insert(w)
			}
			assert.Equal(t, tt.want, trie.StartsWith(tt.args.prefix))
		})
	}
}
