package main

import (
	"fmt"
	"os"
)

var Indices map[int]int

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "The program accepts only one argument")
		return
	}
	sourceCode := os.Args[1]
	Indices = make(map[int]int)
	indices(sourceCode)
	parse(sourceCode)
}

func parse(sourceCode string) {
	var result string
	ptr := 0
	bytes := make([]byte, 2048)
	for i := 0; i < len(sourceCode); i++ {
		if sourceCode[i] == '>' {
			if ptr == 2047 {
				fmt.Fprintf(os.Stderr, "pointer out of range [%d] with length %d\n", i+1, 2048)
				return
			}
			ptr++
		} else if sourceCode[i] == '<' {
			if ptr == 0 {
				fmt.Fprintln(os.Stderr, "pointer out of range [-1]")
				return
			}
			ptr--
		} else if sourceCode[i] == '+' {
			if bytes[ptr] == 255 {
				fmt.Fprintf(os.Stderr, "byte overflow at pointer %d\n", ptr)
				return
			}
			bytes[ptr]++
		} else if sourceCode[i] == '-' {
			if bytes[ptr] == 0 {
				fmt.Fprintf(os.Stderr, "byte underflow at pointer %d\n", ptr)
				return
			}
			bytes[ptr]--
		} else if sourceCode[i] == '[' {
			if bytes[ptr] == 0 {
				i = Indices[i]
			}
		} else if sourceCode[i] == ']' {
			if bytes[ptr] != 0 {
				i = Indices[i]
			}
		} else if sourceCode[i] == '.' {
			result += string(bytes[ptr])
		}
	}
	fmt.Print(result)
}

func indices(sourceCode string) {
	stack := []int{}
	for i, char := range sourceCode {
		if char == '[' {
			stack = append(stack, i)
		} else if char == ']' {
			Indices[stack[len(stack)-1]] = i
			Indices[i] = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
}
