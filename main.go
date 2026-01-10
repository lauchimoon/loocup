package main

import (
    "fmt"
    "strings"
    "os"
    "path"

    "github.com/lauchimoon/loocup/lexer"
    "github.com/lauchimoon/loocup/token"
    "github.com/lauchimoon/loocup/parser"
    "github.com/lauchimoon/loocup/function"
)

type Result struct {
    Func  function.Function
    LineAt int
}

const (
    PROGRAM_NAME = "loocup"
)

func main() {
    // TODO: use golang flag package
    if len(os.Args) < 3 {
        usage()
        return
    }

    signature := os.Args[1]
    filePath := os.Args[2]

    if !checkExtension(filePath) {
        fmt.Printf("%s: '%s' is not a C file (requires .c or .h extension)\n", PROGRAM_NAME, filePath)
        return
    }

    program, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Printf("%s: %v\n", PROGRAM_NAME, err)
        return
    }

    // TODO: We use this to show it exactly as it's written. It may change
    programSplit := strings.Split(string(program), "\n")
    target := function.MakeFromSignature(signature)
    declarations := CollectFunctionDeclarations(string(program))
    for _, decl := range declarations {
        if decl.Func.MatchesCriteria(target) {
            line := decl.LineAt - 1
            fmt.Printf("%s:%d: %s\n", filePath, decl.LineAt, programSplit[line])
        }
    }
}

func usage() {
    fmt.Printf("usage: %s <signature> <file>\n", PROGRAM_NAME)
    fmt.Printf("  <signature> looks like type(arg1, arg2, ..., argN)\n")
    fmt.Printf("  <file> is a .c file or .h file\n")
}

func checkExtension(s string) bool {
    ext := path.Ext(s)
    return ext == ".c" || ext == ".h"
}

func CollectFunctionDeclarations(program string) []Result {
    results := []Result{}
    tokens := lexer.LexerMake(program).Lex()
    i := 0
    readingMultilineComment := false

    for i < len(tokens) {
        if tokens[i].Kind == token.OPEN_MULTICOMMENT {
            readingMultilineComment = true
        }

        for i < len(tokens) && readingMultilineComment {
            if tokens[i].Kind == token.CLOSE_MULTICOMMENT {
                readingMultilineComment = false
            }

            i++
        }

        isFuncDecl, semicolonIndex := parser.IsFunctionDeclaration(tokens, i)
        if isFuncDecl {
            f := function.MakeFromDeclarationTokens(tokens[i:semicolonIndex+1])
            results = append(results, Result{f, tokens[i].Line})
            i = semicolonIndex
        }

        i++
    }

    return results
}

