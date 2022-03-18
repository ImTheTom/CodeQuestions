package main

import "fmt"

const (
	ONES = iota + 1
	TENTHS
	HUNDREDTHS
	THOUSANDTHS
)

func DoQuestion(first int) string {
	var total string
	place := ONES
	for first > 0 {
		current := first % 10
		first = first / 10
		switch place {
		case ONES:
			total = fmt.Sprintf("%s%s", convertToStringOnes(current), total)
		case TENTHS:
			total = fmt.Sprintf("%s%s", convertToStringTenths(current), total)
		case HUNDREDTHS:
			total = fmt.Sprintf("%s%s", convertToStringHundreds(current), total)
		case THOUSANDTHS:
			total = fmt.Sprintf("%s%s", convertToStringThousands(current), total)
		}

		place++
	}
	return total
}

func convertToStringOnes(v int) string {
	switch v {
	case 0:
		return ""
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	}
	return ""
}

func convertToStringTenths(v int) string {
	switch v {
	case 0:
		return ""
	case 1:
		return "X"
	case 2:
		return "XX"
	case 3:
		return "XXX"
	case 4:
		return "XL"
	case 5:
		return "L"
	case 6:
		return "LX"
	case 7:
		return "LXX"
	case 8:
		return "LXXX"
	case 9:
		return "XC"
	}
	return ""
}

func convertToStringHundreds(v int) string {
	switch v {
	case 0:
		return ""
	case 1:
		return "C"
	case 2:
		return "CC"
	case 3:
		return "CCC"
	case 4:
		return "CD"
	case 5:
		return "D"
	case 6:
		return "DC"
	case 7:
		return "DCC"
	case 8:
		return "DCCC"
	case 9:
		return "CM"
	}
	return ""
}

func convertToStringThousands(v int) string {
	switch v {
	case 0:
		return ""
	case 1:
		return "M"
	case 2:
		return "MM"
	case 3:
		return "MMM"
	}
	return ""
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion(1)
	fmt.Printf("Got %v\n", result)
}
