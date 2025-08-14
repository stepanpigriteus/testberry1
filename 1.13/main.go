package main

import "fmt"

func main() {
	a := 1
	b := 2

	a = a + b
	b = a - b //a уже обновилось = 3, а еще b = 2
	a = a - b // b обновило значение  = 1 и  a еще равно  3

	//XOR смотрел, но не принцип 100% не воспроизведу
	// a ^= b
	// b ^= a
	// a ^= b

	fmt.Println(a, b)
}
