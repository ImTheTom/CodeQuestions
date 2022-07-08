package main

func Rotate(arr []int, pos int) {
	n := len(arr)
	tmpArr := make([]int, n)

	k := 0

	for i := pos; i < n; i++ {
		tmpArr[k] = arr[i]
		k++
	}

	for i := 0; i < pos; i++ {
		tmpArr[k] = arr[i]
		k++
	}

	for i, v := range tmpArr {
		arr[i] = v
	}
}
