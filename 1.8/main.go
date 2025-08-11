package main

import "log"

func main() {
	OnBite(10, 1, 200)
}

func OnBite(position int, change int, num int64) {

	if position < 0 || position > 63 {
		log.Println("incorrect position")
		return
	}

	switch change {
	case 0:
		num = num &^ (1 << position)
		log.Println(num)
	case 1:
		num = num | (1 << position)
		log.Println(num)
	default:
		log.Println("incorrect variant")
	}

}
