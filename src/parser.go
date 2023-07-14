package main

import (
	"fmt"
	"strings"
)

func Parse(input string) []string {
    res := make([]string, 0)
    input = strings.ReplaceAll(input, " ", "")
    var curToken string
    for i, char := range input {
        if char == ' ' {
            continue
        }
        if (i != 0) && char == '(' {
            lastChar := input[i-1]
            if lastChar == ')' || (!isOperator(string(lastChar)) && lastChar != '(') {
                res = append(res, "*")
            }
        }
        if isOperator(string(char)) || isParenthesis(string(char)) {
            if curToken != "" {
                res = append(res, curToken)
            }
            res = append(res, string(char))
            curToken = ""
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

    stack := Stack[string]()
    stack.Push("(")

    for _, token := range infix {
        if token == "(" {
            stack.Push(token)
        } else if token == ")" {
            for stack.Top() != "(" {
                res = append(res, stack.Pop())
            }
            stack.Pop()
        } else if isOperator(token) {
            for precedence[stack.Top()] > precedence[token] {
                res = append(res, stack.Pop())
            }
            stack.Push(token)
        } else if !isOperator(token) && !isParenthesis(token) {
            res = append(res, token)
        } else {
            fmt.Printf("INVALID TOKEN %v\n", token)
        }
    }

    return res
}

func isOperator(c string) bool {
    return c == "*" || c == "+" || c == "-" || c == "/" || c == "^"
}

func isParenthesis(c string) bool {
    return c == "(" || c == ")"
}
