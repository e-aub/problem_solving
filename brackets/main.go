package main

import (
	"fmt"
	"os"
)

// map values that are opening brackets to  keys that are closing brackets
var Brackets = map[rune]rune{')': '(', '}': '{', ']': '['}

func clean(text string) string {
	var result string
	// iterate over text
	for _, character := range text {
		// Add character to the result if it is a bracket
		if character == '(' || character == '[' || character == '{' || character == ')' || character == ']' || character == '}' {
			result += string(character)
		}
	}
	return result
}

func brackets(arg string) string {
	stack := []rune{}
	// Iterate over brackets
	for _, bracket := range arg {
		// Check if the bracket is a closing bracket and save its opening
		if opening, isClosing := Brackets[bracket]; isClosing {
			// Check if the closing bracket has its opening in the last element in the stack
			// If it doesnt exist return "Error"
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return "Error"
			}
			// If it exists then remove it
			stack = stack[:len(stack)-1]
		} else {
			// Append opening brackets to the stack
			stack = append(stack, bracket)
		}
	}
	// If the stack is not empty that means that an opening bracket doesent have its pair
	// Return "Error"
	if len(stack) != 0 {
		return "Error"
	}
	// Otherwise return "Ok", because all opening brackets found its closing
	return "Ok"
}

func main() {
	// Remove program name from arguments
	args := os.Args[1:]
	// Check if the number of arguments is valid
	if len(args) >= 1 {
		// Iterate over arguments clean each arg from everything except brackets using clean() function
		// Then check with brackets function if each bracket has its closing
		for _, arg := range args {
			// Print the result returned by brackets
			fmt.Println(brackets(clean(arg)))
		}
		return
	}
	// Print usage error if there's no arguments
	fmt.Println("Enter at least one argument")
}
