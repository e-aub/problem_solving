# [Brackets](https://github.com/01-edu/public/tree/master/subjects/brackets) from [01 Edu System](https://github.com/01-edu/public/tree/master/)

## Description

This Go program validates if a series of input strings are correctly bracketed. It processes each input string to ensure that every opening bracket has a corresponding and correctly positioned closing bracket. The program recognizes three types of brackets:

- Parentheses: `()`
- Square brackets: `[]`
- Curly braces: `{}`

Symbols other than brackets are ignored in the validation process. If an argument is correctly bracketed, the program outputs `OK`; otherwise, it outputs `Error`. A string with no brackets is considered correctly bracketed and will output `OK`.

## Usage

To use the program, run it with one or more string arguments. Each argument should be enclosed in quotes. The program will process each argument independently and print the result for each one.

### Example Commands

```bash
# Valid brackets
$ go run . '(johndoe)' | cat -e
OK$

# Invalid brackets
$ go run . '([)]' | cat -e
Error$

# Multiple arguments, some valid and some invalid
$ go run . '' '{[(0 + 0)(1 + 1)](3*(-1)){()}}' | cat -e
OK$
OK$

# If no arguments are provided, the program will print
$ go run . | cat -e
Enter at least one argument$
