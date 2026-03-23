package main

import "fmt"

func DoQuestion(words []string) int {
	max := 0
	for i := 0; i < len(words)-1; i++ {
		currRes := make([]int, len(words))
		for j := i + 1; j < len(words); j++ {
			if DoesNotShareSameCharacter(words[i], words[j]) {
				currRes[j] = len(words[i]) * len(words[j])
			}
		}

		tmpMax := MaxValue(currRes)

		if tmpMax > max {
			max = tmpMax
		}
	}
	return max
}

func DoesNotShareSameCharacter(first, second string) bool {
	resMap := make(map[rune]int)

	for _, v := range first {
		resMap[v]++
	}

	for _, v := range second {
		if resMap[v] > 0 {
			return false
		}
	}

	return true
}

func MaxValue(values []int) int {
	max := -1
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]string{
		"abcw", "baz", "foo", "bar", "xtfn", "abcdef",
	})
	fmt.Printf("Got %v\n", result)
}
