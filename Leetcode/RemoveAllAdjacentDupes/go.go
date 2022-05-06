package main

import "fmt"

func DoQuestion(s string, k int) string {
	for i := 0; i < len(s); i++ {
		if i < 0 {
			i = 0
		}
		if len(s) == 0 {
			break
		}
		currLen := len(s)
		curr := s[i]
		same := true
		totalInARow := 1

		for j := i + 1; j < i+k && j < currLen; j++ {
			if curr != s[j] {
				same = false
				break
			}
			totalInARow++
		}

		if !same {
			continue
		}

		if totalInARow != k {
			continue
		}

		firstHalf := s[0:i]
		secondHalf := s[i+k : currLen]
		s = firstHalf + secondHalf
		i -= k
	}
	return s
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion("deeedbbcccbdaa", 3)
	fmt.Printf("Got %v\n", result)
}
