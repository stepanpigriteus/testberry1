package main

import "fmt"

func main() {
	s := "лаврыба"
	fmt.Println(swapString(s))

}

func swapString(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}
