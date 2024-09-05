package main

import (
	"unicode"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
	"github.com/01-edu/z01"
)

func main() {
	table := [10]byte{}

	for j := 0; j < 5; j++ {
		for i := 0; i < random.IntBetween(7, 10); i++ {
			table[i] = byte(random.IntBetween(13, 126))
		}
		challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table)
	}
	table2 := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table2)
}

func PrintMemory(arr [10]byte) {
	base := "0123456789abcdef"

	var final []string

	for _, decimal := range arr {
		var temp string
		var res []int

		for decimal > 0 {
			remainder := decimal % 16
			res = append(res, int(remainder))
			decimal /= 16
		}
		for j := len(res) - 1; j >= 0; j-- {
			temp += string(base[res[j]])
		}
		final = append(final, temp)
	}
	for i, res := range final {
		if res == "" {
			Printer("00")
		} else if len(res) == 1 {
			Printer("0" + res)
		} else {
			Printer(res)
		}
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else if (i+1)%4 != 0 && i != len(final)-1 {
			z01.PrintRune(' ')
		}
		if i == len(final)-1 {
			z01.PrintRune('\n')
		}
	}
	for _, letter := range arr {
		if unicode.IsGraphic(rune(letter)) && letter != ' ' {
			z01.PrintRune(rune(letter))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func Printer(s string) {
	for _, let := range s {
		z01.PrintRune(let)
	}
}

// func main() {
// 	arr := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
// 	base := "0123456789abcdef"

// 	var final []string

// 	for _, decimal := range arr {
// 		var temp string
// 		var res []int

// 		for decimal > 0 {
// 			remainder := decimal % 16
// 			res = append(res, int(remainder))
// 			decimal /= 16
// 		}
// 		for j := len(res) - 1; j >= 0; j-- {
// 			temp += string(base[res[j]])
// 		}
// 		final = append(final, temp)
// 	}
// 	for i, res := range final {
// 		if res == "" {
// 			Printer("00")
// 		} else {
// 			Printer(res)
// 		}
// 		if (i+1)%4 == 0 {
// 			z01.PrintRune('\n')
// 		} else if (i+1)%4 != 0 && i != len(final)-1 {
// 			z01.PrintRune(' ')
// 		}
// 		if i == len(final)-1 {
// 			z01.PrintRune('\n')
// 		}
// 	}
// 	for _, letter := range arr {
// 		if unicode.IsGraphic(rune(letter)) {
// 			z01.PrintRune(rune(letter))
// 		} else {
// 			z01.PrintRune('.')
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func Printer(s string) {
// 	for _, let := range s {
// 		z01.PrintRune(let)
// 	}
// }
