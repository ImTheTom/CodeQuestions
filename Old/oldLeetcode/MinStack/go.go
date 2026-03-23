package main

import "fmt"

const MinDefault = 2147483647

type MinStack struct {
	Values []int
	Min    int
}

func Constructor() MinStack {
	return MinStack{
		Values: []int{},
		Min:    MinDefault,
	}
}

func (this *MinStack) Push(val int) {
	if val < this.Min {
		this.Min = val
	}
	this.Values = append(this.Values, val)
}

func (this *MinStack) Pop() {
	lastIndex := len(this.Values) - 1
	poppedValue := this.Values[lastIndex]
	this.Values = this.Values[0:lastIndex]

	if poppedValue == this.Min {
		this.recalculateMin()
	}
}

func (this *MinStack) Top() int {
	lastIndex := len(this.Values) - 1
	return this.Values[lastIndex]
}

func (this *MinStack) GetMin() int {
	return this.Min
}

func (this *MinStack) recalculateMin() {
	tmpMin := MinDefault
	for _, v := range this.Values {
		if v < tmpMin {
			tmpMin = v
		}
	}
	this.Min = tmpMin
}

func main() {
	fmt.Println("Running main")
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Printf("Got getMin %v\n", obj.GetMin())
	obj.Pop()
	fmt.Printf("Got getMin %v\n", obj.Top())
	fmt.Printf("Got getMin %v\n", obj.GetMin())
}
