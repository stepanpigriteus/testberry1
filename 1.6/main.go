package main

import (
	"context"
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	// выберите вариант остановки:
	// 1. канал остановки
	// 2. Закрытие канала
	// 3. context.Context
	// 4. Таймер time.After
	// 5.  Атомарный флаг
	// 6. Return
	// 7. runtime.Goexit
	// 8. panic
	runKillGo(7)
}

func runKillGo(mode int) {
	switch mode {
	case 1:
		stop := make(chan struct{})
		go func() {
			for {
				select {
				case <-stop:
					log.Println("Stop signal received")
					return
				default:
					log.Println("Working...")
					time.Sleep(300 * time.Millisecond)
				}
			}
		}()
		time.Sleep(1 * time.Second)
		close(stop)

	case 2:
		data := make(chan int)
		go func() {
			for val := range data {
				log.Println("Got:", val)
			}
			log.Println("Channel closed, stopping")
		}()
		data <- 1
		data <- 2
		close(data)

	case 3:
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			for {
				select {
				case <-ctx.Done():
					log.Println("Context canceled")
					return
				default:
					log.Println("Working...")
					time.Sleep(300 * time.Millisecond)
				}
			}
		}()
		time.Sleep(1 * time.Second)
		cancel()

	case 4:
		go func() {
			timer := time.After(1 * time.Second)
			for {
				select {
				case <-timer:
					log.Println("Time is up!")
					return
				default:
					log.Println("Working...")
					time.Sleep(300 * time.Millisecond)
				}
			}
		}()

	case 5:
		var stop int32
		go func() {
			for {
				if atomic.LoadInt32(&stop) == 1 {
					log.Println("Stopped by flag")
					return
				}
				log.Println("Working...")
				time.Sleep(300 * time.Millisecond)
			}
		}()
		time.Sleep(1 * time.Second)
		atomic.StoreInt32(&stop, 1)

	case 6:
		go func() {
			for i := 0; i < 3; i++ {
				log.Println("Step", i)
				time.Sleep(100 * time.Millisecond)
			}
			log.Println("Done")
		}()

	case 7:
		go func() {
			defer log.Println("Cleanup before exit")
			log.Println("Working...")
			runtime.Goexit()
			log.Println("Never printed")
		}()

	case 8:
		go func() {
			defer func() { recover() }()
			log.Println("Working...")
			panic("Stop goroutine")
		}()

	default:
		log.Println("Остальных вариантов пока не знаю)")
		return
	}

	time.Sleep(2 * time.Second)
}
