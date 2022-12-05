package main

import (
	"fmt"
	"time"
)

// const file = "text.txt"

const file = "actual.txt"

var scoreing = map[byte]int{
	65:  27,
	66:  28,
	67:  29,
	68:  30,
	69:  31,
	70:  32,
	71:  33,
	72:  34,
	73:  35,
	74:  36,
	75:  37,
	76:  38,
	77:  39,
	78:  40,
	79:  41,
	80:  42,
	81:  43,
	82:  44,
	83:  45,
	84:  46,
	85:  47,
	86:  48,
	87:  49,
	88:  50,
	89:  51,
	90:  52,
	97:  1,
	98:  2,
	99:  3,
	100: 4,
	101: 5,
	102: 6,
	103: 7,
	104: 8,
	105: 9,
	106: 10,
	107: 11,
	108: 12,
	109: 13,
	110: 14,
	111: 15,
	112: 16,
	113: 17,
	114: 18,
	115: 19,
	116: 20,
	117: 21,
	118: 22,
	119: 23,
	120: 24,
	121: 25,
	122: 26,
}

func DoQuestion(input []string) int {
	total := 0

	for _, line := range input {
		dup := findDuplicateByte(line)

		total += scoreing[dup]
	}

	return total
}

func findDuplicateByte(line string) byte {
	firstCompartment, secondCompartment := make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(line)/2; i++ {
		firstCompartment[line[i]]++
	}

	for i := len(line) / 2; i < len(line); i++ {
		secondCompartment[line[i]]++
	}

	for item := range firstCompartment {
		if secondCompartment[item] != 0 {
			return item
		}
	}

	// Sake of simplicity, panicing here
	panic("idk")
}

func main() {
	start := time.Now()

	lines, err := readFile(file)
	if err != nil {
		fmt.Println("reading failed " + err.Error())
	}

	fmt.Printf("loaded file with %d lines\n", len(lines))

	fmt.Println("Running main")
	result := DoQuestion(lines)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
