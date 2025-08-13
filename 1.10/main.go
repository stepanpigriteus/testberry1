package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var temper []float32 = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0.0, 2, 4.4, -3, -0}
	fmt.Println(temperAggr(temper))
}

func temperAggr(arr []float32) map[int][]float32 {
	result := make(map[int][]float32)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, r := range arr {
		wg.Add(1)
		go func(r float32) {
			defer wg.Done()
			key := findKey(r)
			mu.Lock()
			result[key] = append(result[key], r)
			mu.Unlock()
		}(r)
	}
	wg.Wait()
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
