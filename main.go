package main

import (
    "fmt"
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
    f := FunctionMake("int", "f", []FuncArg{
        FuncArgMake("int", "a"),
        FuncArgMake("int", "b"),
    })
    g := FunctionMake("bool", "g", []FuncArg{
        FuncArgMake("int", "a"),
        FuncArgMake("int", "b"),
    })

    fmt.Println(f)
    fmt.Println(g)
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

func (f Function) String() string {
    s := f.RetType + " " + f.Name + "("
    for i, arg := range f.Args {
        if i + 1 >= len(f.Args) {
            s += arg.Type + " " + arg.Name
        } else {
            s += arg.Type + " " + arg.Name + ", "
        }
    }

    return s + ")"
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
