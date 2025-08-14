package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.
// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func main() {
	var arr []string
	var res []string
	countWords := 10000000
	for i := 0; i < countWords; i++ {
		arr = append(arr, wordGen())
	}

	res = setInMap(arr)

	fmt.Println("Result len:", len(res))

	//Тесты
	// start := time.Now()
	// _ = setInMap(arr)
	// elapsed := time.Since(start)
	// fmt.Println("Длинна:", len(arr))
	// fmt.Println("\nМасссив: \nВремя выполнения: ", elapsed)
	// startq := time.Now()
	// _ = setInMapConc(arr)
	// elapsedq := time.Since(startq)
	// fmt.Println("Длинна:", len(arr))
	// fmt.Println("\nВремя выполнения: ", elapsedq)

}

func wordGen() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	length := rand.Intn(8) + 3
	word := make([]rune, length)
	for i := range word {
		word[i] = letters[rand.Intn(len(letters))]
	}
	return string(word)
}

func setInMap(arr []string) []string {
	mapa := make(map[string]struct{})
	var res []string
	if len(arr) == 0 {
		return res
	}
	for _, r := range arr {
		if _, ok := mapa[r]; ok {
			continue
		}
		mapa[r] = struct{}{}
	}
	for key := range mapa {
		res = append(res, key)
	}

	return res
}

// / Логично, что выйгрыша от этой функции не получил))
func setInMapConc(arr []string) []string {
	var res []string
	var wg sync.WaitGroup
	var mu sync.Mutex
	mapa := make(map[string]struct{})
	if len(arr) == 0 {
		return res
	}
	workersNum := runtime.NumCPU()
	chunkSize := (len(arr) + workersNum - 1) / workersNum

	for i := 0; i < workersNum; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(arr) {
			end = len(arr)
		}
		if start >= len(arr) {
			break
		}
		wg.Add(1)
		go func(chunk []string) {
			defer wg.Done()
			for _, r := range chunk {
				mu.Lock()
				mapa[r] = struct{}{}
				mu.Unlock()
			}
		}(arr[start:end])

	}
	wg.Wait()
	return res
}
