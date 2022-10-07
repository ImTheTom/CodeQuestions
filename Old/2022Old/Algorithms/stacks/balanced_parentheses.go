package main

const (
	parenthesesOpen  = "("
	parenthesesClose = ")"

	squareOpen  = "["
	squareClose = "]"

	braceOpen  = "{"
	braceClose = "}"
)

func AreBracketsBalanced(exp string) bool {
	stk := sStack{}

	for _, v := range exp {
		vString := string(v)

		if stk.IsEmpty() {
			stk.Push(vString)
		} else if stk.Peek() == parenthesesOpen && vString == parenthesesClose {
			stk.Pop()
		} else if stk.Peek() == squareOpen && vString == squareClose {
			stk.Pop()
		} else if stk.Peek() == braceOpen && vString == braceClose {
			stk.Pop()
		} else {
			stk.Push(vString)
		}
	}

	return stk.IsEmpty()
}
