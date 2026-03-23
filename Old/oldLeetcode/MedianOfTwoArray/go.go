package main

import (
	"fmt"
	"math"
)

const (
	expected1 = 2.5
	expected2 = 2.0
	expected3 = 1.0
)

var (
	nums11     = []int{1, 2}
	nums12     = []int{3, 4}
	nums21     = []int{1, 3}
	nums22     = []int{2}
	nums31     = []int{1}
	nums32     = []int{}
	binaryFind = []int{1, 2, 2, 6, 7, 9}
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for _, v := range nums2 {
		index := binarySearchForValue(nums1, len(nums1), v)
		nums1 = insert(nums1, index+1, v)
	}

	return medianValue(nums1)
}

// Created from pseudocode https://en.wikipedia.org/wiki/Binary_search_algorithm
func binarySearchForValue(A []int, n, T int) int {
	L := 0
	R := n - 1
	m := int(math.Floor(float64(L+R) / 2))
	for L <= R {
		if A[m] < T {
			L = m + 1
		} else if A[m] > T {
			R = m - 1
		} else {
			return m
		}
		m = int(math.Floor(float64(L+R) / 2))
	}
	return m
}

// https://stackoverflow.com/questions/46128016/insert-a-value-in-a-slice-at-a-given-index
// answered May 15 '20 at 15:03 wasmup
func insert(a []int, index, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func medianValue(A []int) float64 {
	arrayLength := len(A)
	m := int(math.Floor(float64(arrayLength) / 2))
	if arrayLength%2 != 0 {
		return float64(A[m])
	} else {
		L := A[m-1]
		R := A[m]
		return float64(L+R) / 2
	}
}

func main() {
	fmt.Println("Running main")
	index := binarySearchForValue(binaryFind, len(binaryFind), 2)
	fmt.Printf("Got index %d\n", index)
	out := findMedianSortedArrays(nums11, nums12)
	fmt.Printf("Got %f expected %f\n", out, expected1)
	out = findMedianSortedArrays(nums21, nums22)
	fmt.Printf("Got %f expected %f\n", out, expected2)
	out = findMedianSortedArrays(nums31, nums32)
	fmt.Printf("Got %f expected %f\n", out, expected3)
}
