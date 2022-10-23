package main

func ShowcaseRotate() {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7,
	}
	printArray(arr, ArrayBeforePrefix)
	arr = Rotate(arr, 2)
	printArray(arr, ArrayAfterPrefix)
}

func Rotate(nums []int, rotationIndex int) []int {
	sliceLength := len(nums)
	rotatedSlice := make([]int, sliceLength)

	currentIndex := 0
	for i := rotationIndex; i < sliceLength; i++ {
		rotatedSlice[currentIndex] = nums[i]
		currentIndex++
	}

	for i := 0; i < rotationIndex; i++ {
		rotatedSlice[currentIndex] = nums[i]
		currentIndex++
	}

	// Could potentially copy the contents back into the original array.
	// If that's what it wanted.

	return rotatedSlice
}
