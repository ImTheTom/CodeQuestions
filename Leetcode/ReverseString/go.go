package main

import "fmt"

func DoQuestion(s []byte) {
	swapPos := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		tmpByte := s[swapPos]
		s[swapPos] = s[i]
		s[i] = tmpByte
		swapPos -= 1
	}
}

func main() {
	fmt.Println("Running main")
	param := []byte{
		'h', 'e', 'l', 'l', 'o',
	}
	DoQuestion(param)
	fmt.Printf("Got %v\n", param)
}
