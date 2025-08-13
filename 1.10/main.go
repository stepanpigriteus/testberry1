package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	n := rand.Intn(100000000)
	temper := make([]float32, n)
	for i := 0; i < n; i++ {
		temper[i] = rand.Float32()*200 - 100
	}
	start := time.Now()
	_ = temperAggrSeq(temper)
	elapsedSeq := time.Since(start)

	start = time.Now()
	_ = temperAggr(temper)
	elapsedConc := time.Since(start)

	fmt.Printf("Тупое: %v\n", elapsedSeq)
	fmt.Printf("Конкур с чанками: %v\n", elapsedConc)
	fmt.Printf("Разница в сторону конкура: %.2fx\n", float64(elapsedSeq)/float64(elapsedConc))
}

func temperAggr(arr []float32) map[int][]float32 {
	numWorkers := runtime.NumCPU()
	chunkSize := len(arr) / numWorkers
	if chunkSize == 0 {
		return temperAggrSeq(arr)
	}

	result := make(map[int][]float32)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = len(arr)
		}

		wg.Add(1)
		go func(chunk []float32) {
			defer wg.Done()
			localResult := make(map[int][]float32)
			for _, r := range chunk {
				key := findKey(r)
				localResult[key] = append(localResult[key], r)
			}
			mu.Lock()
			for key, values := range localResult {
				result[key] = append(result[key], values...)
			}
			mu.Unlock()
		}(arr[start:end])
	}

	wg.Wait()
	return result
}

// первый последовательнй и кривой вариант
func temperAggrSeq(arr []float32) map[int][]float32 {
	result := make(map[int][]float32)
	if len(arr) == 0 {
		fmt.Println("Empty arr!")
	}
	for _, r := range arr {
		if r == 0 {
			result[0] = append(result[0], 0.0)
			continue
		}
		g := math.Abs(float64(r))

		if g > 0 && g < 10 {
			result[0] = append(result[0], r)
		}
		if r < 0 && r > -10.0 {
			result[0] = append(result[0], r)
			continue
		}
		firstDigit := findKey(r)
		result[firstDigit] = append(result[firstDigit], r)
	}

	return result
}

func findKey(num float32) int {
	g := math.Abs(float64(num))
	if num == 0 || (g > 0 && g < 10) {
		return 0
	}
	absVal := int(math.Abs(float64(num)))          //сначала в инт
	digits := int(math.Log10(float64(absVal))) + 1 //старший разряд
	step := int(math.Pow(10, float64(digits-1)))   //десятки
	res := (int(num) / step) * step                //восстанавливаем десятки
	return res

}
