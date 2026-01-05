package main

import (
    "fmt"

    "github.com/lauchimoon/loocup/lexer"
//    "github.com/lauchimoon/loocup/parser"
    "github.com/lauchimoon/loocup/function"
)

const (
    PROGRAM_NAME = "loocup"
)

func main() {
    f := function.MakeFromTokens(lexer.LexerMake("int add(int a, int b);").Lex())
    g := function.MakeFromSignature("bool(int, int)")

    fmt.Println("f:", f)
    fmt.Println("g:", g)
    fmt.Println("f matches g criteria?", f.MatchesCriteria(g))
}
