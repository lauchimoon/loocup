package function

import (
    "slices"
    "strings"

    "github.com/lauchimoon/loocup/parser"
    "github.com/lauchimoon/loocup/token"
)

type Function struct {
    RetType string
    Name    string
    Args    []FuncArg
}

func Make(retType, name string, args []FuncArg) Function {
    return Function{
        RetType: retType,
        Name: name,
        Args: args,
    }
}

func MakeFromSignature(sig string) Function {
    openParenIndex := strings.IndexByte(sig, '(')
    retType := sig[:openParenIndex]
    return Make(retType, "x", getArgs(sig))
}

func MakeFromTokens(tokens []token.Token) Function {
    isFuncDecl, semicolonIndex := parser.IsFunctionDeclaration(tokens, 0)
    if !isFuncDecl {
        return Function{"", "", []FuncArg{}}
    }

    openParenIndex := slices.Index[[]token.Token, token.Token](tokens, token.Token{
        Kind: token.TOKEN_OPEN_PAREN, Value: "(",
    })

    // TODO: consider modifiers like unsigned, signed, pointers...
    retType := tokens[0].Value
    name := tokens[openParenIndex - 1].Value
    args := ""
    argList := tokens[openParenIndex + 1:semicolonIndex]
    i := 0
    for i < len(argList) - 1 {
        args += argList[i].Value + " "
        i++
    }

    fString := retType + "(" + args + ")"
    f := MakeFromSignature(fString)
    f.Name = name

    return f
}

func getArgs(sig string) []FuncArg {
    openParenIndex := strings.IndexByte(sig, '(')
    closeParenIndex := strings.IndexByte(sig, ')')
    argsSig := sig[openParenIndex+1:closeParenIndex]

    args := []FuncArg{}
    for _, arg := range strings.Split(argsSig, ",") {
        fields := strings.Fields(strings.Trim(arg, " "))
        if len(fields) < 1 {
            continue
        }

        argType := fields[0]
        var argName string

        if len(fields) >= 2 {
            argName = fields[1]
        }
        args = append(args, FuncArgMake(argType, argName))
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
