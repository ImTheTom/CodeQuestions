package main

import "fmt"

func DoQuestion(s string, t string) bool {
	s = stripHashtags(s)
	t = stripHashtags(t)
	if s == t {
		return true
	}
	return false
}

func stripHashtags(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			end := len(s)
			if i == 0 {
				s = s[i+1 : end]
				i--
			} else {
				s = s[0:i-1] + s[i+1:end]
				i -= 2
			}
		}
	}
	return s
}

func main() {
	fmt.Println("Running main")
	result := stripHashtags("#a#c")
	fmt.Printf("Got %v\n", result)
}
