package main

import (
    "fmt"
    "os"
    "bufio"

    "calculator/parser"
    "calculator/eval"
)

func main() {
    fmt.Println("Welcome to the calculator!")
    scanner := bufio.NewScanner(os.Stdin)
    var input string
    var postfix []string
    var infix []string
    var res string
    for {
        fmt.Print("> ")
        if !scanner.Scan() {
            continue
        }
        input = scanner.Text()
        infix = parser.Parse(input)
        postfix = parser.InfixToPostfix(infix)
        res = eval.EvaluatePosfix(postfix)
        fmt.Println(res)
    }
}
