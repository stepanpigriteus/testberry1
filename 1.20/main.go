package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "lorem ipsum color dolor sit amet"
	fmt.Println(reverseOrder(s))
}

func reverseOrder(s string) string {
	arr := strings.Fields(s)
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return strings.Join(arr, " ")
}
