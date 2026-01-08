package parser

import (
    "slices"

    "github.com/lauchimoon/loocup/token"
)

func IsFunctionDeclaration(tokens []token.Token, i int) (bool, int) {
    idx := i
    if idx >= len(tokens) || (tokens[idx].Kind != token.SYMBOL && tokens[idx].Kind != token.KEYWORD) {
        return false, -1
    }

    if tokens[idx].Kind == token.KEYWORD && (!isTypeSpec(tokens[idx]) && tokens[idx].Value != "void") {
        return false, -1
    }

    // We checked the type properly
    idx++

    // In case we have pointers or something else
    if tokens[idx].Kind != token.SYMBOL {
        for idx < len(tokens) && tokens[idx].Kind != token.SYMBOL {
            // We found semicolon before another symbol or opening parenthesis.
            if tokens[idx].Kind == token.SEMICOLON {
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

    if (tokens[idx].Kind != token.OPEN_PAREN) {
        return false, -1
    }

    // Parse arguments
    idx++

    if !isWord(tokens[idx]) && tokens[idx].Kind != token.CLOSE_PAREN {
        return false, -1
    }

    for tokens[idx].Kind != token.CLOSE_PAREN {
        argStart := idx

        for idx < len(tokens) && tokens[idx].Kind != token.COMMA && tokens[idx].Kind != token.CLOSE_PAREN {
            idx++
        }

        args := tokens[argStart:idx]
        if !isValidArgs(args) {
            return false, -1
        }

        if idx < len(tokens) && tokens[idx].Kind == token.COMMA {
            idx++
            if tokens[idx].Kind == token.CLOSE_PAREN {
                return false, -1
            }
        }
    }

    idx++
    if idx >= len(tokens) || tokens[idx].Kind != token.SEMICOLON {
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
    return t.Kind == token.SYMBOL || (t.Kind == token.KEYWORD && (isTypeSpec(t) || t.Value == "void"))
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
