package main

import "fmt"

func ShowcaseArrays() {
	fmt.Println("Building array")
	exampleArray := [5]int{3, 2, 1, 4, 5}
	traversal(exampleArray)
	newArr := insertion(9, exampleArray)
	removedArr := delete(2, newArr)
	traversal(removedArr)
	fmt.Printf("search result was %d\n", search(27, removedArr))
	fmt.Printf("search result was %d\n", search(9, removedArr))
	sortedArr := sort(removedArr)
	traversal(sortedArr)
}

func traversal(nums [5]int) {
	fmt.Println("Traversing array")
	for i, v := range nums {
		fmt.Printf("Index - %d, Value - %d\n", i, v)
	}
}

func insertion(newValue int, nums [5]int) [6]int {
	fmt.Printf("Inserting value %d\n", newValue)
	var newArr = [6]int{nums[0], nums[1], nums[2], nums[3], nums[4], newValue}
	return newArr
}

func delete(index int, nums [6]int) [5]int {
	fmt.Printf("Deleting value at index %d\n", index)
	var newArr = [5]int{}
	currentIndex := 0
	for i := 0; i < len(nums); i++ {
		if i == index {
			continue
		}
		newArr[currentIndex] = nums[i]
		currentIndex++
	}
	return newArr
}

func search(value int, nums [5]int) int {
	fmt.Printf("Searching for value %d\n", value)
	index := -1
	for i, v := range nums {
		if v == value {
			return i
		}
	}
	return index
}

func sort(nums [5]int) [5]int {
	fmt.Println("Sorting array")
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
