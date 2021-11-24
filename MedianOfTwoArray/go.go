package main

import "fmt"

const expected = 2.5

var (
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
}

func main() {
	fmt.Println("Running main")
	out := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("Got %f expected %f\n", out, expected)
}
