package main

import (
	"fmt"
	"time"
)

type TimeMap struct {
	dict map[string][]data
}

type data struct {
	timestamp int
	value     string
}

func Constructor() TimeMap {
	return TimeMap{
		dict: make(map[string][]data),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.dict[key] = append(this.dict[key], data{
		timestamp: timestamp,
		value:     value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	var (
		result     = ""
		dataValues = this.dict[key]

		left  = 0
		right = len(dataValues) - 1
	)

	for left <= right {
		mid := (left + right) / 2

		if dataValues[mid].timestamp <= timestamp {
			result = dataValues[mid].value
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	tm := Constructor()
	tm.Set("foo", "bar", 1)
	fmt.Println(tm.Get("foo", 1))
	fmt.Println(tm.Get("foo", 3))

	tm.Set("foo", "bar2", 4)
	fmt.Println(tm.Get("foo", 4))
	fmt.Println(tm.Get("foo", 5))

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
