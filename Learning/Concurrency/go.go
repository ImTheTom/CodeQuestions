package main

import (
	"fmt"
	"sync"
	"time"
)

const NumberToSquare = 5000

const NumberOfWorkers = 5

type Job struct {
	ID    int
	Value int
}

func CreateInputArray() []int {
	originals := make([]int, NumberToSquare)

	for i := 0; i < NumberToSquare; i++ {
		originals[i] = i
	}

	return originals
}

func Synchronous(digits []int) []int {
	squared := make([]int, NumberToSquare)

	for i := 0; i < NumberToSquare; i++ {
		squared[i] = i * i
		time.Sleep(100 * time.Microsecond)
	}

	return squared
}

func BasicWorkers(digits []int) []int {
	squared := make([]int, NumberToSquare)

	var (
		wg         sync.WaitGroup
		jobChannel = make(chan Job)
	)

	wg.Add(NumberOfWorkers)
	for i := 0; i < NumberOfWorkers; i++ {
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
		time.Sleep(100 * time.Microsecond)
	}
}

func ActualWorker(digits []int) []int {
	jobs := make(chan int, len(digits))
	results := make(chan int, len(digits))

	squared := make([]int, NumberToSquare)

	for w := 0; w < NumberOfWorkers; w++ {
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
	}
}

func main() {
	fmt.Println("Running main")

	originals := CreateInputArray()

	BasicWorkers(originals)

	// fmt.Printf("Got %v\n", result)
}
