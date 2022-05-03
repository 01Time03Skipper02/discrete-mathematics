package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func makeTokens(str string) []string {
	var (
		position = 0
		tokens   = make([]string, 0, 0)
	)
	for position < len(str) && str[position] != ')' {
		//fmt.Println(str[position])
		if str[position] == ' ' {
			position++
		}
		if str[position] == '-' || str[position] == '*' || str[position] == '+' || unicode.IsDigit(rune(str[position])) {
			tokens = append(tokens, str[position:position+1])
			position++
		} else {
			var (
				position2 int
				flag      int = 1
			)
			for i := position + 1; i < len(str); i++ {
				if str[i] == '(' {
					flag++
				}
				if str[i] == ')' {
					flag--
				}
				if flag == 0 {
					position2 = i + 1
					break
				}
			}
			tokens = append(tokens, str[position:position2])
			position = position2
		}
	}
	return tokens
}

func parse(str []string) []string {
	var (
		res []string
	)
	res = append(res, "(")
	for i := 0; i < len(str); i++ {
		fmt.Println(res)
		if len(str[i]) > 1 {
			buf := str[i][1 : len(str[i])-1]
			buf2 := makeTokens(buf)
			buf3 := parse(buf2)
			//fmt.Println(buf3)
			for j := 0; j < len(buf3); j++ {
				//fmt.Println(buf3[j])
				res = append(res, buf3[j])
			}
		} else {
			if str[i] == "-" || str[i] == "+" || str[i] == "*" {
				if len(res) == 0 {
					res = append(res, str[i])
				} else {
					res = append(res, "")
					for j := len(res) - 1; j > 1; j-- {
						res[j] = res[j-1]
					}
					res[1] = str[i]
				}
			}
			runes := []rune(str[i])
			if unicode.IsDigit(runes[0]) && !(str[i] == "-" || str[i] == "+" || str[i] == "*") && len(runes) == 1 {
				res = append(res, str[i])
			}
		}
	}
	res = append(res, ")")
	return res
}

func calculate(expression []string) int {
	var (
		res int
	)

	return res
}

func main() {
	var (
		expression string
	)
	expression, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	expression = expression[:len(expression)-2]
	//fmt.Println(makeTokens(expression))        //check making array of tokens
	fmt.Println(parse(makeTokens(expression))) //check parse

}
