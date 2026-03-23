package main

import "fmt"

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	value := this.queue[len(this.queue)-1]
	this.queue = this.queue[0 : len(this.queue)-1]
	return value
}

func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func main() {
	fmt.Println("Running main")
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.Top())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Empty())
}
