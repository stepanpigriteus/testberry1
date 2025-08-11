package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	concMap(3)
}

func concMap(n int) {
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	mapa := map[int]int{}
	stop := make(chan struct{})

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
					key := rand.Intn(1000000)
					val := rand.Intn(100)
					mutex.Lock()
					mapa[key] = val
					mutex.Unlock()
				}
				time.Sleep(50 * time.Millisecond)
			}
		}()
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)
			mutex.RLock()
			fmt.Println("len:", len(mapa))
			mutex.RUnlock()
		}
	}()

	go func() {
		<-time.After(5 * time.Second)
		close(stop)
	}()

	wg.Wait()
}
