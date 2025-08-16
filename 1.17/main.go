package main

import "fmt"

func main() {
	size := 1000000
	seek := 921
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	binary(arr, seek)
}

func binary(arr []int, g int) int {
	left, right := 0, len(arr)-1
	counter := 0
	for left <= right {
		counter++

		mid := (left + right) / 2
		if arr[mid] == g {
			fmt.Println("шагов:", counter)
			return mid
		} else if arr[mid] < g {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
