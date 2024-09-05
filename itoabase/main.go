package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	for i := 0; i < 30; i++ {
		value := random.IntBetween(-1000000, 1000000)
		base := random.IntBetween(2, 16)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, value, base)
	}
	for i := 0; i < 5; i++ {
		base := random.IntBetween(2, 16)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, random.MaxInt, base)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, random.MinInt, base)
	}
}

func ItoaBase(value, base int) string {
	hex := "0123456789ABCDEF"

	stringBase := hex[:base]
	var result string
	sign := ""
	if value == -9223372036854775808 {
		maxInt := uint(9223372036854775808)
		uintBase := uint(base)
		for maxInt > 0 {
			result = string(stringBase[maxInt%uintBase]) + result
			maxInt /= uintBase
		}
		return "-" + result
	} else if value < 0 {
		sign = "-"
		value = -value
	}

	for value > 0 {
		result = string(stringBase[value%base]) + result
		value /= base
	}
	result = sign + result
	return result
}
