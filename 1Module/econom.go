package main

import (
	"bufio"
	"fmt"
	"os"
)

func makeTokens(str string) []string {
	var (
		position = 0
		tokens   = make([]string, 0, 0)
	)
	for position < len(str) {
		if str[position] == ' ' {
			position++
		} else if str[position] == '(' && position != 0 {
			token := ""
			cnt := 1
			flag := 1
			for flag != 0 {
				if str[position] == '(' && cnt != 1 {
					flag++
					cnt++
				}
				if str[position] == ')' {
					flag--
					cnt++
				}
				token += string(str[position])
				cnt++
				position++
			}
			tokens = append(tokens, token)
		} else {
			tokens = append(tokens, str[position:position+1])
			position++
		}
	}
	return tokens
}

func makeHashTable(expression []string) map[string]bool {
	var (
		hash = make(map[string]bool)
	)
	for i := 0; i < len(expression); i++ {
		if expression[i] == "#" || expression[i] == "$" || expression[i] == "@" {
			key := "(" + expression[i] + " " + expression[i+1] + " " + expression[i+2] + ")"
			hash[key] = true
		}
		if len(expression[i]) > 1 {
			tokens := makeTokens(expression[i])
			newHash := makeHashTable(tokens)
			for k, v := range newHash {
				hash[k] = v
			}
		}
	}
	return hash
}

func main() {
	var (
		expression string
	)
	expression, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	expression = expression[:len(expression)-2]
	//fmt.Println(makeTokens(expression))
	//fmt.Println(makeHashTable(makeTokens(expression)))
	fmt.Println(len(makeHashTable(makeTokens(expression))))
}
