package parser

import (
    "slices"

    "github.com/lauchimoon/loocup/token"
)

func IsFunctionDeclaration(tokens []token.Token, i int) (bool, int) {
    idx := i
    if tokens[idx].Kind != token.TOKEN_SYMBOL && tokens[idx].Kind != token.TOKEN_KEYWORD {
        return false, -1
    }

    // TODO: check for modifiers like unsigned, signed...
    if tokens[idx].Kind == token.TOKEN_KEYWORD && !isPrimitiveType(tokens[idx].Value) {
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
        if tokens[idx].Kind == token.TOKEN_COMMA {
            if tokens[idx + 1].Kind != token.TOKEN_SYMBOL && tokens[idx + 1].Kind != token.TOKEN_KEYWORD {
                return false, -1
            }

            idx++
        }

        if isWord(tokens[idx]) {
            if tokens[idx + 1].Kind == token.TOKEN_COMMA {
                idx++ 
            } else if tokens[idx + 1].Kind == token.TOKEN_CLOSE_PAREN {
                idx++
            } else if isWord(tokens[idx + 1]) {
                if tokens[idx + 1].Kind == token.TOKEN_KEYWORD {
                    return false, -1
                }

                if (idx + 2 >= len(tokens)) || (tokens[idx + 2].Kind != token.TOKEN_COMMA && tokens[idx + 2].Kind != token.TOKEN_CLOSE_PAREN) {
                    return false, -1
                }
                idx += 2
            } else {
                return false, -1
            }
        }
    }

    idx++
    if idx >= len(tokens) || tokens[idx].Kind != token.TOKEN_SEMICOLON {
        return false, -1
    }

    return true, idx
}

func isPrimitiveType(s string) bool {
    types := []string{
        "char", "double", "float",
        "int", "long", "short",
        "void",
    }

    return slices.Index[[]string, string](types, s) != -1
}

func isWord(t token.Token) bool {
    return t.Kind == token.TOKEN_SYMBOL || (t.Kind == token.TOKEN_KEYWORD && isPrimitiveType(t.Value))
}
