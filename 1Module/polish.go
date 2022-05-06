package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func makeTokens(str string) []string {
	var (
		position = 0
		tokens   = make([]string, 0, 0)
	)
	for position < len(str) {
		if str[position] == ' ' {
			position++
		} else {
			tokens = append(tokens, str[position:position+1])
			position++
		}
	}
	return tokens
}

func calculate(expression []string) int {
	var (
		res int
	)
	for i := 0; i < len(expression); i++ {
		if expression[i] == ")" {
			var x int
			a, _ := strconv.Atoi(expression[i-2])
			b, _ := strconv.Atoi(expression[i-1])
			if expression[i-3] == "+" {
				x = a + b
				expression[i] = strconv.Itoa(x)
			}
			if expression[i-3] == "-" {
				x = a - b
				expression[i] = strconv.Itoa(x)
			}
			if expression[i-3] == "*" {
				x = a * b
				expression[i] = strconv.Itoa(x)
			}
			expression = append(expression[:i-4], expression[i:]...)
			i = 0
		}
		fmt.Println(expression)
	}
	res, _ = strconv.Atoi(expression[0])
	return res
}

func main() {
	var (
		expression string
	)
	expression, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	expression = expression[:len(expression)-2]
	fmt.Println(makeTokens(expression))
	fmt.Println(calculate(makeTokens(expression)))

}
