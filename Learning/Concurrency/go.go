package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID    int
	Value int
}

func CreateInputArray() []int {
	originals := make([]int, 100)

	for i := 0; i < 100; i++ {
		originals[i] = i
	}

	return originals
}

func Synchronous(digits []int) []int {
	squared := make([]int, 100)

	for i := 0; i < 100; i++ {
		squared[i] = i * i
		time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
	}

	return squared
}

func BasicWorkers(digits []int) []int {
	squared := make([]int, 100)

	var (
		wg         sync.WaitGroup
		jobChannel = make(chan Job)
	)

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go basicWorker(i, &wg, jobChannel, squared)
	}

	for i := 0; i < len(digits); i++ {
		jobChannel <- Job{
			ID:    digits[i],
			Value: digits[i],
		}
	}

	close(jobChannel)
	wg.Wait()

	return squared
}

func basicWorker(id int, wg *sync.WaitGroup, jobChannel <-chan Job, result []int) {
	defer wg.Done()
	for v := range jobChannel {
		result[v.ID] = v.Value * v.Value
		time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
	}
}

func ActualWorker(digits []int) []int {
	jobs := make(chan int, len(digits))
	results := make(chan int, len(digits))

	squared := make([]int, 100)

	for w := 1; w <= 3; w++ {
		go actualWorker(w, jobs, results)
	}

	for i := 0; i < len(digits); i++ {
		jobs <- digits[i]
	}
	close(jobs)

	for i := 0; i < len(digits); i++ {
		squared[i] = <-results
	}

	return squared
}

func actualWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * j
		time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
	}
}

func main() {
	fmt.Println("Running main")

	originals := CreateInputArray()

	result := ActualWorker(originals)

	fmt.Printf("Got %v\n", result)
}
