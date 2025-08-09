package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// принимает количество воркеров и таймер остановки
	Task(3, 5)
}

func Task(n int, timer int) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timer)*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		ticker := time.NewTicker(30 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				select {
				case ch <- rand.Intn(255):
				case <-ctx.Done():
					close(ch)
					return
				}
			}
		}
	}()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case res, ok := <-ch:
					if !ok {
						return
					}
					fmt.Println("Worker", id, "Res:", res)
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}

	wg.Wait()
}
