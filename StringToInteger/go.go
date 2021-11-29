package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	i = "42"
	e = 42

	i2 = "    042"
	e2 = 42

	i3 = "    -042"
	e3 = -42

	i4 = "00000-42a1234"
	e4 = 0

	i5 = "9223372036854775808"
	e5 = 2147483647

	i6 = "-91283472332"
	e6 = -2147483648

	MIN_INT = -2147483648
	MAX_INT = 2147483647
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	seenSign := false
	firstSeen := false
	isPositive := true
	var value int = 0

	for _, v := range s {
		if v == 43 {
			if seenSign || firstSeen {
				break
			} else {
				seenSign = true
			}
		} else if v == 45 {
			if seenSign || firstSeen {
				break
			} else {
				seenSign = true
				isPositive = false
			}
		} else if v >= 48 && v <= 57 {
			firstSeen = true
			value = value*10 + int(v) - 48
			if value > MAX_INT {
				if !isPositive {
					value = MIN_INT
				} else {
					value = MAX_INT
				}
				break
			}
		} else {
			break
		}
	}

	if !isPositive && value > 0 {
		value = value - (value * 2)
	}

	return value
}

func doFunction(i string, e int) {
	a := myAtoi(i)
	fmt.Printf("For - %s I got %d when I wanted %d\n", i, a, e)
	if a != e {
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Running main")
	doFunction(i, e)
	doFunction(i2, e2)
	doFunction(i3, e3)
	doFunction(i4, e4)
	doFunction(i5, e5)
	doFunction(i6, e6)
}
