package main

import (
    "fmt"

    "github.com/lauchimoon/loocup/lexer"
    "github.com/lauchimoon/loocup/function"
)

const (
    PROGRAM_NAME = "loocup"
)

func main() {
    f := function.MakeFromSignature("int(int, int)")
    g := function.MakeFromSignature("bool(int, int)")

    fmt.Println("f:", f)
    fmt.Println("g:", g)
    fmt.Println("f matches g criteria?", f.MatchesCriteria(g))

    fmt.Println(lexer.LexerMake("int *f(int a, int b) { return a + b; }").Lex())
    fmt.Println(lexer.LexerMake("int *f(int a[12], int b) { return a[0xbeef] * b; }").Lex())
    fmt.Println(lexer.LexerMake("// Hello world").Lex())
    fmt.Println(lexer.LexerMake("/* Hello world */").Lex())
}
