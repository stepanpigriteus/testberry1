package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	results := make([]int, len(arr))

	var wg sync.WaitGroup
	ch := make(chan struct {
		index int
		value int
	}, len(arr))
	for i, r := range arr {
		wg.Add(1)
		go func(i int, r int) {
			defer wg.Done()
			ch <- struct {
				index int
				value int
			}{index: i, value: r * r}
		}(i, r)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		results[res.index] = res.value
	}
	fmt.Println(results)
}
