package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func DoQuestion(s string) bool {
	s = stripNonAlphanumericCharacter(s)
	halfWayPoint := len(s) / 2
	for i := 0; i < halfWayPoint; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func stripNonAlphanumericCharacter(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	stripped := reg.ReplaceAllString(s, "")
	return strings.ToLower(stripped)
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion("A man, a plan, a canal: Panama")
	fmt.Printf("Got %v\n", result)
}
