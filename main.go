package main

import (
    "fmt"
    "strings"

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
    program := `/* This tests function declarations like int this_is_shouldnt_be_read(int a); */
// Here's a single line comment with a declaration: bool dont_read_this(void);
int add(int a, int b);
int sub(int a, int b);
bool is_positive(int x);
void nothing();
size_t len(char *, char *);
int *dir(int a);

int add(int a, int b) { return a + b; }
int sub(int a, int b) { return a + b; }
bool is_positive(int x) { return x > 0; }
void nothing() {}
int *dir(int a) { return &a; }`

    // TODO: We use this to show it exactly as it's written. It may change
    programSplit := strings.Split(program, "\n")
    target := function.MakeFromSignature("int(int, int)")
    declarations := CollectFunctionDeclarations(program)
    for _, decl := range declarations {
        if decl.Func.MatchesCriteria(target) {
            line := decl.LineAt - 1
            fmt.Printf("program:%d: %s\n", line, programSplit[line])
        }
    }
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
