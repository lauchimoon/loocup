package lexer

import (
    "io"
    "strings"
    "unicode"

    "github.com/lauchimoon/loocup/token"
)

type Lexer struct {
    Source    string
    SourceLen int
    Cursor    int
}

func LexerMake(src string) *Lexer {
    return &Lexer{
        Source: src,
        SourceLen: len(src),
        Cursor: 0,
    }
}

// This lexer only focuses on function declarations, so it will read:
// - symbols (types and variable names)
// - punctuation (such as (, ), {, }, ; and ,)
// TODO: ignore function bodies!
func (l *Lexer) Lex() []token.Token {
    tokens := []token.Token{}
    for l.Cursor < l.SourceLen {
        c := l.Current()
        buffer := strings.Builder{}

        // Skip whitespace
        if unicode.IsSpace(rune(c)) {
            l.Advance()
            c = l.Current()
            for unicode.IsSpace(rune(c)) {
                l.Advance()
                c = l.Current()
            }
        }

        // Symbols
        if unicode.IsLetter(rune(c)) {
            buffer.WriteByte(c)
            l.Advance()
            c = l.Current()

            for unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) {
                buffer.WriteByte(c)
                l.Advance()
                c = l.Current()
            }

            tokens = append(tokens, token.Token{
                Kind: token.TOKEN_SYMBOL,
                Value: buffer.String(),
            })
        }

        // Others
        if c == '(' {
            tokens = append(tokens, token.Token{
                Kind: token.TOKEN_OPEN_PAREN,
                Value: "(",
            })
        } else if c == ')' {
            tokens = append(tokens, token.Token{
                Kind: token.TOKEN_CLOSE_PAREN,
                Value: ")",
            })
        } else if c == '{' {
            tokens = append(tokens, token.Token{
                Kind: token.TOKEN_OPEN_BRACKET,
                Value: "{",
            })
        } else if c == '}' {
            tokens = append(tokens, token.Token{
                Kind: token.TOKEN_CLOSE_BRACKET,
                Value: "}",
            })
        } else if c == ';' {
            tokens = append(tokens, token.Token{
                Kind: token.TOKEN_SEMICOLON,
                Value: ";",
            })
        } else if c == ',' {
            tokens = append(tokens, token.Token{
                Kind: token.TOKEN_COMMA,
                Value: ",",
            })
        }

        l.Advance()
    }

    return tokens
}

func (l *Lexer) Advance() {
    if (l.Cursor + 1 > l.SourceLen) {
        return
    }

    l.Cursor++
}

func (l *Lexer) Current() byte {
    return l.Source[l.Cursor]
}

func (l *Lexer) Peek() (byte, error) {
    if (l.Cursor + 1 >= l.SourceLen) {
        return 0, io.EOF
    }

    return l.Source[l.Cursor + 1], nil
}
