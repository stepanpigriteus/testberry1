package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	signOut(3)
}

func signOut(n int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					log.Printf("Worker %d: end(\n", id)
					return
				default:
					log.Printf("Worker %d: work-work\n", id)
					time.Sleep(1 * time.Second)
				}
			}
		}(i + 1)
	}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case sig := <-sigChan:
			log.Println("Получен сигнал:", sig)
			cancel()
			done := make(chan struct{}) // канал заглушка для селекта
			go func() {
				wg.Wait() // ожиидание завершения воркеров
				close(done)
			}()
			select {
			case <-done:
				log.Println("Горутинки завершены")
			case <-time.After(5 * time.Second):
				log.Println("Завершение по таймауту")
			}
			return
		case <-ticker.C:
			log.Println("Still alive!")
		}
	}
}
