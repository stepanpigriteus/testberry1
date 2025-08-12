package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	convee()
}

func convee() {
	nums := []int{10, 20, 30, 40, -50}
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go stageOne(nums, ch1, &wg)
	wg.Add(1)
	go stageTwo(ch1, ch2, &wg)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for result := range ch2 {
			log.Printf("Result: %d\n", result)
		}
	}()

	wg.Wait()
}

func stageOne(nums []int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)
	for _, num := range nums {
		out <- num
		time.Sleep(200 * time.Millisecond)
	}
}

func stageTwo(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)

	for num := range in {
		result := num * 2
		out <- result
		time.Sleep(100 * time.Millisecond)
	}
}
