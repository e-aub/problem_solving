package main

import (
	"fmt"
	"os"
)

var Indices map[int]int

func main() {
	commands := os.Args[1]
	Indices = make(map[int]int)
	indices(commands)
	parse(commands)
}

func parse(commands string) {
	ptr := 0
	bytes := make([]byte, 2048)
	for i := 0; i < len(commands); i++ {
		if commands[i] == '>' {
			ptr++
		} else if commands[i] == '<' {
			ptr--
		} else if commands[i] == '+' {
			bytes[ptr]++
		} else if commands[i] == '-' {
			bytes[ptr]--
		} else if commands[i] == '[' {
			if bytes[ptr] == 0 {
				i = Indices[i]
			}
		} else if commands[i] == ']' {
			if bytes[ptr] != 0 {
				i = Indices[i]
			}
		} else if commands[i] == '.' {
			fmt.Print(string(bytes[ptr]))
		}
	}
}

func indices(commands string) {
	stack := []int{}
	for i, char := range commands {
		if char == '[' {
			stack = append(stack, i)
		} else if char == ']' {
			Indices[stack[len(stack)-1]] = i
			Indices[i] = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	fmt.Println(Indices)
}
