package parser

import (
//    "fmt"
    "slices"

    "github.com/lauchimoon/loocup/token"
)

func IsFunctionDeclaration(tokens []token.Token, i int) (bool, int) {
    idx := i
    if idx >= len(tokens) || (tokens[idx].Kind != token.TOKEN_SYMBOL && tokens[idx].Kind != token.TOKEN_KEYWORD) {
        return false, -1
    }

    if tokens[idx].Kind == token.TOKEN_KEYWORD && (!isTypeSpec(tokens[idx]) && tokens[idx].Value != "void") {
        return false, -1
    }

    // We checked the type properly
    idx++

    // In case we have pointers or something else
    if tokens[idx].Kind != token.TOKEN_SYMBOL {
        for idx < len(tokens) && tokens[idx].Kind != token.TOKEN_SYMBOL {
            // We found semicolon before another symbol or opening parenthesis.
            if tokens[idx].Kind == token.TOKEN_SEMICOLON {
                return false, -1
            }

            idx++
        }
    }

    if idx >= len(tokens) {
        return false, -1
    }

    // Move to '('
    idx++

    if (tokens[idx].Kind != token.TOKEN_OPEN_PAREN) {
        return false, -1
    }

    // Parse arguments
    idx++

    if !isWord(tokens[idx]) && tokens[idx].Kind != token.TOKEN_CLOSE_PAREN {
        return false, -1
    }

    for tokens[idx].Kind != token.TOKEN_CLOSE_PAREN {
        nextIdx := token.FindByKind(tokens[idx:], token.TOKEN_COMMA)
        if nextIdx == -1 {
            nextIdx = token.FindByKind(tokens[idx:], token.TOKEN_CLOSE_PAREN)
            if nextIdx == -1 {
                return false, -1
            }
        }

        args := tokens[idx:idx+nextIdx]
        if !isValidArgs(args) {
            return false, -1
        }

        idx += nextIdx
        if tokens[idx].Kind == token.TOKEN_COMMA {
            idx++
        }
    }

    idx++
    if idx >= len(tokens) || tokens[idx].Kind != token.TOKEN_SEMICOLON {
        return false, -1
    }

    return true, idx
}

func isTypeSpec(t token.Token) bool {
    types := []string{
        "char", "double", "float",
        "int", "long", "short",
        "const", "unsigned", "signed",
    }

    return slices.Index[[]string, string](types, t.Value) != -1
}

func isWord(t token.Token) bool {
    return t.Kind == token.TOKEN_SYMBOL || (t.Kind == token.TOKEN_KEYWORD && (isTypeSpec(t) || t.Value == "void"))
}

func isValidArgs(args []token.Token) bool {
    counts := map[string]int{}
    for _, arg := range args {
        if isTypeSpec(arg) {
            counts[arg.Value]++
        }
    }

    // If we have more than one of these, it's forbidden
    forbidden := []string{
        "char", "short", "int", "signed",
        "unsigned", "float", "double", "void", 

    }

    for _, kind := range forbidden {
        if counts[kind] > 1 {
            if kind == "long" && counts[kind] == 2 {
                continue
            }

            return false
        }
    }

    return true
}
