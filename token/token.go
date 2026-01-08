package token

type TokenKind int
const (
    TOKEN_OPEN_PAREN = iota
    TOKEN_CLOSE_PAREN
    TOKEN_OPEN_CURLY
    TOKEN_CLOSE_CURLY
    TOKEN_OPEN_BRACKET
    TOKEN_CLOSE_BRACKET
    TOKEN_SEMICOLON
    TOKEN_COMMA
    TOKEN_OPERATOR
    TOKEN_NEWLINE
    TOKEN_ASTERISK
    TOKEN_NUMBER
    TOKEN_SYMBOL
    TOKEN_KEYWORD
    TOKEN_COMMENT
    TOKEN_OPEN_MULTICOMMENT
    TOKEN_CLOSE_MULTICOMMENT
    TOKEN_PREPROC
)

type Token struct {
    Kind TokenKind
    Value string
}

func FindByKind(tokens []Token, kind TokenKind) int {
    for i, t := range tokens {
        if t.Kind == kind {
            return i
        }
    }

    return -1
}
