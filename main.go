package main

import (
    "fmt"
    "strings"
)

type FuncArg struct {
    Type string
    Name string
}

type Function struct {
    RetType string
    Name    string
    Args    []FuncArg
}

func main() {
    f := FunctionMakeFromSignature("int(int, int)")
    g := FunctionMakeFromSignature("bool(int, int)")

    fmt.Println("f:", f)
    fmt.Println("g:", g)
    fmt.Println("f matches g criteria?", f.MatchesCriteria(g))
}

func FunctionMake(retType, name string, args []FuncArg) Function {
    return Function{
        RetType: retType,
        Name: name,
        Args: args,
    }
}

func FuncArgMake(typ, name string) FuncArg {
    return FuncArg{
        Type: typ,
        Name: name,
    }
}

func FunctionMakeFromSignature(sig string) Function {
    openParenIndex := strings.IndexByte(sig, '(')
    retType := sig[:openParenIndex]

    return Function{
        RetType: retType,
        Name: "x",
        Args: getArgs(sig),
    }
}

func getArgs(sig string) []FuncArg {
    openParenIndex := strings.IndexByte(sig, '(')
    closeParenIndex := strings.IndexByte(sig, ')')
    argsSig := sig[openParenIndex+1:closeParenIndex]

    args := []FuncArg{}
    for i, typ := range strings.Split(argsSig, ",") {
        typ = strings.Trim(typ, " ")
        args = append(args, FuncArgMake(typ, fmt.Sprintf("x%d", i)))
    }

    return args
}

func (f Function) String() string {
    s := f.RetType + " " + f.Name + "("
    for i, arg := range f.Args {
        if i + 1 >= len(f.Args) {
            s += arg.Type + " " + arg.Name
        } else {
            s += arg.Type + " " + arg.Name + ", "
        }
    }

    return s + ");"
}

// Names don't matter for now. We only care about:
// - matching return type
// - matching number of arguments
// - matching argument types
func (f Function) MatchesCriteria(criteria Function) bool {
    if len(f.Args) != len(criteria.Args) {
        return false
    }

    for i, fArg := range f.Args {
        critArg := criteria.Args[i]
        if fArg.Type != critArg.Type {
            return false
        }
    }

    return f.RetType == criteria.RetType
}
