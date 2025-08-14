package main

import "fmt"

func main() {
	typeOfx("string")
	typeOfx(4)
	typeOfx(false)
}

func typeOfx(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println("int >>>", v)
	case string:
		fmt.Println("String >>>", v)
	case bool:
		fmt.Println("Bool >>>", v)
	default:
		fmt.Println("Другой")
	}
}
