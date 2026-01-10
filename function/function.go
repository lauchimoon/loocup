package function

import (
    "strings"
    "regexp"

    "github.com/lauchimoon/loocup/lexer"
    "github.com/lauchimoon/loocup/token"
)

var (
    identifierMatcher = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
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
    tokens := lexer.LexerMake(sig).Lex()
    openParenIndex := token.FindByKind(tokens, token.OPEN_PAREN)
    closeParenIndex := token.FindByKind(tokens, token.CLOSE_PAREN)
    if openParenIndex == -1 || closeParenIndex == -1 {
        return Function{"", "", []FuncArg{}}
    }

    retType := makeRetType(tokens[:openParenIndex])
    args := makeArgs(tokens[openParenIndex:closeParenIndex+1])
    if args == nil {
        return Function{"", "", []FuncArg{}}
    }

    return Make(strings.Trim(retType, " "), "x", args)
}

func makeRetType(tokens []token.Token) string {
    retType := ""
    for _, t := range tokens {
        retType += t.Value + " "
    }

    return retType
}

func makeArgs(tokens []token.Token) []FuncArg {
    args := []FuncArg{}
    if tokens[0].Kind != token.OPEN_PAREN || tokens[len(tokens)-1].Kind != token.CLOSE_PAREN {
        return nil
    }

    group := []token.Token{}

    for _, t := range tokens[1:len(tokens)] {
        if t.Kind == token.COMMA || t.Kind == token.CLOSE_PAREN {
            arg := buildArg(group)
            if len(arg.Type) == 0 && len(arg.Name) == 0 {
                return []FuncArg{}
            }

            args = append(args, arg)
            group = []token.Token{}
        } else {
            group = append(group, t)
        }
    }

    return args
}

func buildArg(tokens []token.Token) FuncArg {
    if len(tokens) == 0 {
        return FuncArg{}
    }

    lastIdx := len(tokens) - 1
    lastToken := tokens[lastIdx]
    if len(tokens) > 1 && isName(lastToken) {
        return FuncArgMake(
            formatType(tokens[:lastIdx]),
            lastToken.Value,
        )
    }

    typ := formatType(tokens[:lastIdx + 1])
    return FuncArgMake(typ, "")
}

func isName(t token.Token) bool {
    excludeSymbols := "[]*&."
    if strings.ContainsAny(t.Value, excludeSymbols) {
        return false
    }

    return identifierMatcher.MatchString(t.Value)
}

func formatType(tokens []token.Token) string {
    tokensString := []string{}
    for _, t := range tokens {
        tokensString = append(tokensString, t.Value)
    }

    typ := strings.Join(tokensString, " ")
    return strings.Trim(typ, " ")
}

func MakeFromDeclarationTokens(tokens []token.Token) Function {
    openParenIndex := token.FindByKind(tokens, token.OPEN_PAREN)
    retType := ""
    for _, t := range tokens[:openParenIndex - 1] {
        retType += t.Value + " "
    }

    name := tokens[openParenIndex - 1].Value
    args := ""
    semicolonIndex := token.FindByKind(tokens, token.SEMICOLON)
    argList := tokens[openParenIndex + 1:semicolonIndex]
    for _, arg := range argList {
        args += arg.Value + " "
    }

    fString := retType + "(" + args + ")"
    f := MakeFromSignature(fString)
    f.Name = name

    return f
}

// TODO: this looks ugly sometimes... what can we do about it?
func (f Function) String() string {
    s := f.RetType + " " + f.Name + "("
    for i, arg := range f.Args {
        if i + 1 >= len(f.Args) {
            s += arg.Type + formatNameString(arg.Name)
        } else {
            s += arg.Type + formatNameString(arg.Name) + ", "
        }
    }

    return s + ");"
}

func formatNameString(name string) string {
    if len(name) == 0 {
        return ""
    }

    return " " + name
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

    return strings.Contains(f.RetType, criteria.RetType)
}
