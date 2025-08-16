package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	size := 9000000
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100000)
	}
	startq := time.Now()
	quicksortq(arr)
	elapsedq := time.Since(startq)
	fmt.Println("\nВремя выполнения: ", elapsedq)

}

func quicksortq(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	index := len(arr) / 2
	pivot := arr[index]

	var left, right []int
	for i, v := range arr {
		if i == index {
			continue
		}
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quicksortq(left), pivot), quicksortq(right)...)
}
