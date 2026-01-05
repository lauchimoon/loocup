package main

import (
    "fmt"

    "github.com/lauchimoon/loocup/lexer"
    "github.com/lauchimoon/loocup/parser"
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

    fmt.Println("\nTesting some declarations")
    checkDecl("int add(int a, int b);")
    checkDecl("int add(int, int);")
    checkDecl("int add(int a, int);")
    checkDecl("int add(int, int b);")
    checkDecl("int add(int int);")
    checkDecl("int add(int, int)")
    checkDecl("int (int, int);")
    checkDecl("int add(, int);")
    checkDecl("int add(int, int;")
    checkDecl("int addint, int;")
    checkDecl("break add(int a, int b);")
    checkDecl("add(int a, int b);")
}

func checkDecl(src string) {
    tokens := lexer.LexerMake(src).Lex()
    isDecl, _ := parser.IsFunctionDeclaration(tokens, 0)
    fmt.Printf("\"%s\" is a function declaration? %v\n", src, isDecl)
}
