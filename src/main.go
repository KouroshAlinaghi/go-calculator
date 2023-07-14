package main

import (
    "fmt"
    "os"
    "bufio"
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
        infix = Parse(input)
        postfix = InfixToPostfix(infix)
        res = CalculatePosfix(postfix)
        fmt.Println(res)
    }
}
