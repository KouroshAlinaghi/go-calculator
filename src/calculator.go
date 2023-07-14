package main

import (
	"math"
	"strconv"
	"fmt"
)

func CalculatePosfix(postfix []string) string {
    stack := Stack[string]()
    for _, token := range postfix {
        if !isOperator(token) {
            stack.Push(token)
        } else {
            right := stack.Pop()
            left := stack.Pop()
            res := calculateResult(left, token, right)
            stack.Push(res)
        }
    }

    return stack.Top()
}

func calculateResult(left, operator, right string) string {
    var res float64
    l, _ := strconv.ParseFloat(left, 64)
    r, _ := strconv.ParseFloat(right, 64)
    switch operator {
    case "+":
        res = l + r
    case "-":
        res = l - r
    case "*":
        res = l * r
    case "/":
        res = l / r
    case "^":
        res = math.Pow(l, r)
    }
    
    return fmt.Sprintf("%.2f", res)
}
