package token

type Kind int

const (
    OPEN_PAREN = iota
    CLOSE_PAREN
    OPEN_CURLY
    CLOSE_CURLY
    OPEN_BRACKET
    CLOSE_BRACKET
    SEMICOLON
    COMMA
    OPERATOR
    NEWLINE
    ASTERISK
    NUMBER
    SYMBOL
    KEYWORD
    COMMENT
    OPEN_MULTICOMMENT
    CLOSE_MULTICOMMENT
    PREPROC
)

type Token struct {
    Kind Kind
    Value string
}

func FindByKind(tokens []Token, kind Kind) int {
    for i, t := range tokens {
        if t.Kind == kind {
            return i
        }
    }

    return -1
}
