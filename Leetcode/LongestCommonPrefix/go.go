package main

import "fmt"

func DoQuestion(strs []string) string {
	currentPrefix := ""
	for i := 0; i < len(strs[0]); i++ {
		possibleExtraPrefix := strs[0][i]
		actualPrefix := true
		for _, v := range strs {
			if len(v) <= i {
				actualPrefix = false
				break
			}
			if possibleExtraPrefix != v[i] {
				actualPrefix = false
				break
			}
		}
		if actualPrefix {
			currentPrefix = fmt.Sprintf("%s%s", string(currentPrefix), string(possibleExtraPrefix))
		} else {
			break
		}
	}
	return string(currentPrefix)
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]string{
		"flower", "flow", "flight",
	})
	fmt.Printf("Got %v\n", result)
}
