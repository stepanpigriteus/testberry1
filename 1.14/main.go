package main

import (
	"fmt"
	"reflect"
)

func main() {
	typeOfx("string")
	typeOfx(4)
	typeOfx(false)
	var ch chan int
	typeOfx(ch)
}

func typeOfx(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println("int >>>", v)
	case string:
		fmt.Println("String >>>", v)
	case bool:
		fmt.Println("Bool >>>", v)
	}

	r := reflect.TypeOf(x).Kind() == reflect.Chan
	if r {
		fmt.Println("Channel")
	}
}
