package main

import "fmt"

func main() {
	s := "asc qwe rt yuio"
	fmt.Println(unicChar(s))
}

func unicChar(s string) bool {
	flag := true
	if len(s) < 2 {
		return flag
	}
	mapa := make(map[rune]struct{})
	for _, char := range s {
		if char == ' ' {
			continue
		}
		if _, exists := mapa[char]; exists {
			return false
		}
		mapa[char] = struct{}{}
	}
	return flag
}
