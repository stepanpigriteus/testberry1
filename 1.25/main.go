package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		sleep(5)
		fmt.Println("1: Woke up!")
	}()
	go func() {
		defer wg.Done()
		sleep(3)
		fmt.Println("2: Woke up!")
	}()

	wg.Wait()
}

func sleep(d int) {
	timer := time.NewTimer(time.Duration(d) * time.Second)
	<-timer.C
	timer.Stop()
}
