package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	TimeToStop(3)
}

func TimeToStop(n int) {
	var wg sync.WaitGroup
	resultChan := make(chan int)

	timer := time.NewTimer(time.Duration(n) * time.Second)
	defer timer.Stop()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(resultChan)
		i := 0
		for {

			select {
			case <-timer.C:
				log.Println("Timer dead!")
				return
			default:
				select {
				case resultChan <- i:
					i++
				case <-timer.C:
					log.Println("Timer went off while sending")
					return
				}
			}
			time.Sleep(5 * time.Millisecond)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case i, ok := <-resultChan:
				if !ok {
					log.Println("Chanell was closed")
					return
				}
				log.Println("Read", i)
			}
			time.Sleep(5 * time.Millisecond)
		}
	}()
	wg.Wait()
}
