package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// const file = "text.txt"

const file = "actual.txt"

var possibleValues = []rune{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

var re = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

type Action struct {
	Amount int
	From   int
	To     int
}

type ListHelper struct {
	Stacks  [][]rune
	Actions []*Action
}

func (l *ListHelper) doActions() {
	for _, a := range l.Actions {
		// Read the elements
		elements := l.Stacks[a.From][(len(l.Stacks[a.From]) - a.Amount):]

		l.Stacks[a.From] = l.Stacks[a.From][:(len(l.Stacks[a.From]) - a.Amount)]

		l.Stacks[a.To] = append(l.Stacks[a.To], elements...)
	}
}

func DoQuestion(input []string) *ListHelper {
	lh := createListHelper(input)

	lh.doActions()

	return lh
}

func createListHelper(input []string) *ListHelper {
	isStacking := true

	stacks := make([][]rune, 12)

	actions := make([]*Action, 0)

	listHelper := &ListHelper{
		Stacks:  stacks,
		Actions: actions,
	}

	for _, value := range input {
		if isStacking {
			if !strings.Contains(value, "[") {
				isStacking = false
				continue
			}

			for i, v := range value {
				if isRuneIn(v, possibleValues) {
					listHelper.Stacks[i/4] = append(listHelper.Stacks[i/4], v)
				}
			}
			continue
		}

		if value == "" {
			continue
		}

		submatchall := re.FindAllString(value, -1)

		amount, _ := strconv.Atoi(submatchall[0])
		from, _ := strconv.Atoi(submatchall[1])
		to, _ := strconv.Atoi(submatchall[2])

		listHelper.Actions = append(listHelper.Actions, &Action{
			Amount: amount,
			From:   from - 1,
			To:     to - 1,
		})
	}

	for i := range listHelper.Stacks {
		listHelper.Stacks[i] = reverseRuns(listHelper.Stacks[i])
	}

	return listHelper
}

func isRuneIn(needle rune, haystack []rune) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func reverseRuns(run []rune) []rune {
	if len(run) == 0 {
		return run
	}

	return append(reverseRuns(run[1:]), run[0])
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
