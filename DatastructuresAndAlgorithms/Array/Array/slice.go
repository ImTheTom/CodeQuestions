package main

import "fmt"

func ShowcaseSlices() {
	fmt.Println("Building Slice")
	exampleSlice := []int{3, 2, 1, 4, 5}
	traversalSlice(exampleSlice)
	newSlice := insertionSlice(9, exampleSlice)
	removedSlice := deleteSlice(2, newSlice)
	traversalSlice(removedSlice)
	fmt.Printf("search result was %d\n", searchSlice(27, removedSlice))
	fmt.Printf("search result was %d\n", searchSlice(9, removedSlice))
	sortedSlice := sortSlice(removedSlice)
	traversalSlice(sortedSlice)
}

func traversalSlice(nums []int) {
	fmt.Println("Traversing slice")
	for i, v := range nums {
		fmt.Printf("Index - %d, Value - %d\n", i, v)
	}
}

func insertionSlice(newValue int, nums []int) []int {
	fmt.Printf("Inserting value in slice %d\n", newValue)
	nums = append(nums, newValue)
	return nums
}

func deleteSlice(index int, nums []int) []int {
	fmt.Printf("Deleting value at index for slice %d\n", index)
	return append(nums[:index], nums[index+1:]...)
}

func searchSlice(value int, nums []int) int {
	fmt.Printf("Searching for value in slice %d\n", value)
	index := -1
	for i, v := range nums {
		if v == value {
			return i
		}
	}
	return index
}

func sortSlice(nums []int) []int {
	fmt.Println("Sorting slice")
	var temp int
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				temp = nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			}
		}
	}
	return nums
}
