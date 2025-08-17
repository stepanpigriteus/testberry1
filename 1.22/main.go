package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).
// Комментарий: в Go тип int справится с такими числами, но обратите внимание на возможное переполнение для ещё больших значений. Для очень больших чисел можно использовать math/big.

func main() {
	a := big.NewInt(1234567890123456789)
	b := big.NewInt(987654321987654321)
	combain(a, b)
}

func combain(a, b *big.Int) big.Int {
	sum := new(big.Int).Add(a, b)
	fmt.Println("Sum:", sum)

	diff := new(big.Int).Sub(a, b)
	fmt.Println("Diff:", diff)

	product := new(big.Int).Mul(a, b)
	fmt.Println("Prod:", product)

	aFloat := new(big.Float).SetInt(a)
	bFloat := new(big.Float).SetInt(b)
	quotient := new(big.Float).Quo(aFloat, bFloat)

	fmt.Println("Quotient:", quotient)

	return *a
}
