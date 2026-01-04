package token

import (
    "fmt"
)

type TokenKind int
const (
    TOKEN_OPEN_PAREN = iota
    TOKEN_CLOSE_PAREN
    TOKEN_OPEN_BRACKET
    TOKEN_CLOSE_BRACKET
    TOKEN_SEMICOLON
    TOKEN_COMMA
    TOKEN_SYMBOL
)

type Token struct {
    Kind TokenKind
    Value string
}

func (t Token) String() string {
    return fmt.Sprintf("%s: %s", kindString(t.Kind), t.Value)
}

func kindString(kind TokenKind) string {
    switch kind {
        case TOKEN_OPEN_PAREN: return "open paren"
        case TOKEN_CLOSE_PAREN: return  "close paren"
        case TOKEN_OPEN_BRACKET: return "open bracket"
        case TOKEN_CLOSE_BRACKET: return "close bracket"
        case TOKEN_SEMICOLON: return "semicolon"
        case TOKEN_COMMA: return "comma"
        case TOKEN_SYMBOL: return "symbol"
    }
    return "?"
}
