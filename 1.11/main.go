package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

func main() {
	a := make([]int, rand.Intn(11000000))
	for i := range a {
		a[i] = rand.Intn(100000000000000)
	}

	b := make([]int, rand.Intn(51000000))
	for i := range b {
		b[i] = rand.Intn(1000000000000000)
	}
	if len(a) < runtime.NumCPU() || len(a) < runtime.NumCPU() || len(a)+len(b) < 800 {
		cross(a, b)
	} else {
		crossWorkerPool(a, b)
	}
	//Тайминги
	start := time.Now()
	_ = cross(a, b)
	elapsed := time.Since(start)
	fmt.Println("Длинна:", len(a), len(b), "\nВ байтах:", len(a)*int(unsafe.Sizeof(a[0])), len(a)*int(unsafe.Sizeof(b[0])))
	fmt.Println("\nБез конкур: \nВремя выполнения: ", elapsed)
	startq := time.Now()
	_ = crossWorkerPool(a, b)
	elapsedq := time.Since(startq)
	fmt.Println("\nКонкур: \nВремя выполнения: ", elapsedq)

}

func cross(a, b []int) []int {
	hash := make(map[int]int)
	var result []int
	for _, r := range a {
		hash[r]++
	}
	for _, v := range b {
		if hash[v] > 0 {
			result = append(result, v)
			hash[v]--
		}
	}
	return result
}

func crossWorkerPool(a, b []int) []int {
	var result []int
	if len(a) == 0 || len(b) == 0 {
		return result
	}
	hash := make(map[int]int, len(a))
	var mu sync.RWMutex

	for _, v := range a {
		hash[v]++
	}
	var wg sync.WaitGroup
	workersNum := runtime.NumCPU()
	results := make([][]int, workersNum)
	chunkSize := (len(b) + workersNum - 1) / workersNum //попыка разбивки на равные части, кроме последней
	for i := 0; i < workersNum; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(b) {
			end = len(b)
		}
		if start >= len(b) {
			break
		}

		wg.Add(1)
		go func(part []int, i int) {
			defer wg.Done()
			localResult := make([]int, 0, len(part)/10)
			localHash := make(map[int]int, len(a))
			mu.RLock()
			for k, v := range hash {
				localHash[k] = v
			}
			mu.RUnlock()
			for _, val := range part {
				if localHash[val] > 0 {
					localResult = append(localResult, val)
					localHash[val]--
				}
			}
			results[i] = localResult
		}(b[start:end], i)

	}
	for _, r := range results {
		result = append(result, r...)
	}

	return result
}
