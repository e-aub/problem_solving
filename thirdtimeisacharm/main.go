package main

import "fmt"

func ThirdTimeIsACharm(str string) string {
	if str == "" || len(str) < 3 {
		return "\n"
	}
	var result string
	for i := 0; i < len(str); i++ {
		if (i+1)%3 == 0 {
			result += string(str[i])
		}
	}

	return result + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
