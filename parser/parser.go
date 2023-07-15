package parser

import (
	"fmt"
	"strings"

	"calculator/stack"
)

func Parse(input string) []string {
    res := make([]string, 0)
    input = strings.ReplaceAll(input, " ", "")
    var curToken string
    for i, char := range input {
        if IsOperator(string(char)) || isParenthesis(string(char)) {
            if char == '-' && (i == 0 || !isDigit(string(input[i-1]))) {
                curToken = "-"
                continue
            }
            if curToken != "" {
                res = append(res, curToken)
            }
            res = append(res, string(char))
            curToken = ""

            if i != 0 && char == '(' {
                lastChar := input[i-1]
                if lastChar == ')' || isDigit(string(lastChar)) {
                    res = append(res, "*")
                }
            }

        } else if i == len(input) - 1 {
            curToken += string(char)
            res = append(res, curToken)
        } else {
            curToken += string(char)
        }
    }

    return res
}

var precedence = map[string]int{
    "(": 1,
    "+": 2,
    "-": 2,
    "*": 3,
    "/": 3,
    "^": 4,
}

func InfixToPostfix(infix []string) []string {
    res := make([]string, 0)

    infix = append(infix, ")")

    stack := stack.Stack[string]()
    stack.Push("(")

    for _, token := range infix {
        if token == "(" {
            stack.Push(token)
        } else if token == ")" {
            for stack.Top() != "(" {
                res = append(res, stack.Pop())
            }
            stack.Pop()
        } else if IsOperator(token) {
            for precedence[stack.Top()] > precedence[token] {
                res = append(res, stack.Pop())
            }
            stack.Push(token)
        } else if !IsOperator(token) && !isParenthesis(token) {
            res = append(res, token)
        } else {
            fmt.Printf("INVALID TOKEN %v\n", token)
        }
    }

    return res
}

func IsOperator(c string) bool {
    return c == "*" || c == "+" || c == "-" || c == "/" || c == "^"
}

func isParenthesis(c string) bool {
    return c == "(" || c == ")"
}

func isDigit(c string) bool {
    return c == "0" ||
           c == "1" ||
           c == "2" ||
           c == "3" ||
           c == "4" ||
           c == "5" ||
           c == "6" ||
           c == "7" ||
           c == "8" ||
           c == "9"
}
