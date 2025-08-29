package main

import (
	"fmt"
	"math/big"
)

// * / + - a b  > 2^20

func calculation(a, b *big.Int, operation string) *big.Int {
	switch {
	case operation == "*":
		return new(big.Int).Mul(a, b)
	case operation == "/":
		return new(big.Int).Div(a, b)
	case operation == "+":
		return new(big.Int).Add(a, b)
	case operation == "-":
		return new(big.Int).Sub(a, b)
	default:
		return big.NewInt(0)
	}
}

func main() {
	a := new(big.Int)
	a.SetString("123123123123123123321321", 10)
	b := new(big.Int)
	b.SetString("321321321321321321123123", 10)
	result := calculation(a, b, "/")
	fmt.Println(result.String())
}
