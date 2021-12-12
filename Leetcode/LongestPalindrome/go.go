package main

import (
	"fmt"
	"math"
)

const (
	s1      = "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	output1 = "bab"
	s2      = "cbbd"
	output2 = "bb"
)

// Still fails long strings for time restrictions TODO

func longestPalindrome(s string) string {
	longestString := ""
	currentString := ""
	knownPalidromes := make(map[string]int)

	for i := 0; i < len(s); i++ {
		currentString = ""
		for j := i; j < len(s); j++ {
			currentString += string(s[j])
			if ok := isPalindrome(currentString, knownPalidromes); ok {
				if len(currentString) > len(longestString) {
					longestString = currentString
				}
			}
		}
	}

	return longestString
}

func isPalindrome(s string, knownPalidromes map[string]int) bool {
	lengthString := float64(len(s))
	loopLength := math.Ceil(lengthString / 2)
	currentString := s

	for i := 0; i < int(loopLength); i++ {
		if knownPalidromes[currentString] == 1 {
			break
		}

		if s[i] != s[len(s)-1-i] {
			return false
		}

		if len(currentString) > 2 {
			currentString = currentString[1 : len(currentString)-1]
		}
	}
	knownPalidromes[s] = 1
	return true
}

func main() {
	fmt.Println("Running main")

	actual1 := longestPalindrome(s1)
	fmt.Printf("For %s Expected %s got %s\n", s1, output1, actual1)
	actual2 := longestPalindrome(s2)
	fmt.Printf("For %s Expected %s got %s\n", s2, output2, actual2)
}
