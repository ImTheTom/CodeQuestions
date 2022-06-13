package main

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
)

// see https://go.dev/doc/tutorial/fuzz

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func main() {
	start := time.Now()

	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
