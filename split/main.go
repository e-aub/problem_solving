package main

import "fmt"

func CustomSplit(str string, separator string) []string {
	start := 0
	var result []string
	for i := 0; i <= len(str)-len(separator); i++ {
		if str[i:i+len(separator)] == separator {
			if toAppend := str[start:i]; toAppend != "" {
				result = append(result, toAppend)
				start = i + len(separator)
			}
			i += len(separator) - 1
		}
	}
	if start < len(str) {
		result = append(result, str[start:])
	}
	return result
}

func main() {
	result := CustomSplit("ab  ff  ff  c  ", "  ")
	for _, word := range result {
		fmt.Println(word)
	}
}
