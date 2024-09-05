package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	for i := 0; i < 50; i++ {
		challenge.Function("Itoa", Itoa, solutions.Itoa, random.Int())
	}
}

func Itoa(n int) string {
	var sign string
	var result string
	if n == 0 {
		return "0"
	}
	if n < 0 {
		sign = "-"
		n *= -1
	}
	for n > 0 {
		result = string((n%10)+48) + result
		n /= 10
	}
	return sign + result
}
