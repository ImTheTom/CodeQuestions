package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// const file = "text.txt"

const file = "actual.txt"

const (
	FirstElf  = 0
	SecondElf = 1

	StartPos = 0
	EndPos   = 1
)

type ElfAssignment struct {
	Start int
	End   int
}

func DoQuestion(input []string) int {
	assignments := make([][]*ElfAssignment, 0)

	for _, value := range input {
		assignment := GenerateElfAssignments(value)

		assignments = append(assignments, assignment)
	}

	total := 0

	for _, assignment := range assignments {
		if DoTheyEncompassEachOther(assignment[FirstElf], assignment[SecondElf]) {
			total++
		}
	}

	return total
}

func GenerateElfAssignments(value string) []*ElfAssignment {
	elfs := strings.Split(value, ",")

	first := GenerateElfAssignment(elfs[FirstElf])
	second := GenerateElfAssignment(elfs[SecondElf])

	return []*ElfAssignment{
		first, second,
	}
}

func GenerateElfAssignment(assignment string) *ElfAssignment {
	selections := strings.Split(assignment, "-")

	start, err := strconv.Atoi(selections[StartPos])
	if err != nil {
		panic(err)
	}

	end, err := strconv.Atoi(selections[EndPos])
	if err != nil {
		panic(err)
	}

	return &ElfAssignment{
		Start: start,
		End:   end,
	}
}

func DoTheyEncompassEachOther(first, second *ElfAssignment) bool {
	if first.Start <= second.Start && first.End >= second.End {
		return true
	}

	if second.Start <= first.Start && second.End >= first.End {
		return true
	}

	return false
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
