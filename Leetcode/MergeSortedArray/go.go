package main

import "fmt"

func DoQuestion(nums1 []int, m int, nums2 []int, n int) {
	lastInserted := 0
	totalInserted := 0
	for i := 0; i < n; i++ {
		for j := lastInserted; j < m; j++ {
			// insert condition
			if nums2[i] < nums1[j] {
				ShiftDownArray(nums1, j, nums2[i])
				m++
				lastInserted = j
				totalInserted++
				break
			}
		}
	}
	for i := totalInserted; i < n; i++ {
		nums1[m] = nums2[totalInserted]
		m++
		totalInserted++
	}
}

func ShiftDownArray(nums1 []int, index, value int) {
	last := nums1[index]
	for i := index + 1; i < len(nums1); i++ {
		tmp := nums1[i]
		nums1[i] = last
		last = tmp
	}
	nums1[index] = value
}

func main() {
	fmt.Println("Running main")
	nums := []int{
		1, 2, 3, 0, 0, 0,
	}
	m := 3
	nums2 := []int{
		2, 5, 6,
	}
	n := 3
	DoQuestion(nums, m, nums2, n)
	fmt.Printf("Got %v\n", nums)
}
