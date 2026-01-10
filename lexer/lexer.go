package lexer

import (
    "io"
    "strings"
    "slices"
    "unicode"

    "github.com/lauchimoon/loocup/token"
)

var (
    keywords = []string{
        "auto", "break", "case", "char",
        "const", "continue", "default", "do",
        "double", "else", "enum", "extern",
        "float", "for", "goto", "if",
        "int", "long", "register", "return",
        "short", "signed", "sizeof", "static",
        "struct", "switch", "typedef", "union",
        "unsigned", "void", "volatile", "while",
    }

    lineCount = 1
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

// We lex everything relevant, the parser takes care of it later
func (l *Lexer) Lex() []token.Token {
    tokens := []token.Token{}
    for l.Cursor < l.SourceLen {
        c := l.Current()
        buffer := strings.Builder{}

        // Skip whitespace
        if c != '\n' && unicode.IsSpace(rune(c)) {
            l.Advance()
            c = l.Current()
            for unicode.IsSpace(rune(c)) {
                l.Advance()
                c = l.Current()
            }
        }

        // Numbers
        if unicode.IsDigit(rune(c)) {
            buffer.WriteByte(c)
            l.Advance()
            c = l.Current()

            hex := (unicode.ToLower(rune(c)) == 'x')
            if (unicode.ToLower(rune(c)) == 'x' ||
                unicode.ToLower(rune(c)) == 'b') {
                    buffer.WriteByte(c)
                    l.Advance()
                    c = l.Current()
            }

            for unicode.IsDigit(rune(c)) || (hex && c >= 'a' && c <= 'f') {
                buffer.WriteByte(c)
                l.Advance()
                c = l.Current()
            }

            tokens = append(tokens, token.Token{
                Kind: token.NUMBER,
                Value: buffer.String(),
                Line: lineCount,
            })
        }

        // Symbols / keywords
        if unicode.IsLetter(rune(c)) || c == '_' {
            buffer.WriteByte(c)
            l.Advance()
            c = l.Current()

            for unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) || c == '_' {
                buffer.WriteByte(c)
                l.Advance()
                c = l.Current()
            }

            value := buffer.String()
            var kind token.Kind = token.SYMBOL
            if isKeyword(value) {
                kind = token.KEYWORD
            }

            tokens = append(tokens, token.Token{
                Kind: kind,
                Value: value,
                Line: lineCount,
            })
        }

        // Multi-lineCount comments or asterisk
        if c == '*' {
            nextC, _ := l.Peek()
            if nextC == '/' {
                l.Advance()
                l.Advance()
                tokens = append(tokens, token.Token{
                    Kind: token.CLOSE_MULTICOMMENT,
                    Value: "*/",
                    Line: lineCount,
                })
            } else {
                tokens = append(tokens, token.Token{
                    Kind: token.ASTERISK,
                    Value: "*",
                    Line: lineCount,
                })
            }
        }

        // Comments
        if c == '/' {
            nextC, _ := l.Peek()
            if nextC == '/' {
                l.Advance()
                l.Advance()
                tokens = append(tokens, token.Token{
                    Kind: token.COMMENT,
                    Value: "//",
                    Line: lineCount,
                })

                c = l.Current()
                for c != '\n' {
                    buffer.WriteByte(c)
                    l.Advance()
                    if l.Cursor >= l.SourceLen {
                        break
                    }
                    c = l.Current()
                }

                tokens = append(tokens, token.Token{
                    Kind: token.COMMENT,
                    Value: buffer.String(),
                    Line: lineCount,
                })
            } else if nextC == '*' {
                l.Advance()
                l.Advance()
                tokens = append(tokens, token.Token{
                    Kind: token.OPEN_MULTICOMMENT,
                    Value: "/*",
                    Line: lineCount,
                })

                c = l.Current()
            }
        }

        // Preprocessor directives
        // TODO: support multi-lineCount preprocessor directives (macros...)
        if c == '#' {
            l.Advance()
            tokens = append(tokens, token.Token{
                Kind: token.PREPROC,
                Value: "#",
                Line: lineCount,
            })

            c = l.Current()
            for c != '\n' {
                buffer.WriteByte(c)
                l.Advance()
                if l.Cursor >= l.SourceLen {
                    break
                }
                c = l.Current()
            }

            tokens = append(tokens, token.Token{
                Kind: token.PREPROC,
                Value: buffer.String(),
                Line: lineCount,
            })
        }

        // Others
        if c == '(' {
            tokens = append(tokens, token.Token{
                Kind: token.OPEN_PAREN,
                Value: "(",
                Line: lineCount,
            })
        } else if c == ')' {
            tokens = append(tokens, token.Token{
                Kind: token.CLOSE_PAREN,
                Value: ")",
                Line: lineCount,
            })
        } else if c == '{' {
            tokens = append(tokens, token.Token{
                Kind: token.OPEN_CURLY,
                Value: "{",
                Line: lineCount,
            })
        } else if c == '}' {
            tokens = append(tokens, token.Token{
                Kind: token.CLOSE_CURLY,
                Value: "}",
                Line: lineCount,
            })
        } else if c == '[' {
            tokens = append(tokens, token.Token{
                Kind: token.OPEN_BRACKET,
                Value: "[",
                Line: lineCount,
            })
        } else if c == ']' {
            tokens = append(tokens, token.Token{
                Kind: token.CLOSE_BRACKET,
                Value: "]",
                Line: lineCount,
            })
        } else if c == ';' {
            tokens = append(tokens, token.Token{
                Kind: token.SEMICOLON,
                Value: ";",
                Line: lineCount,
            })
        } else if c == ',' {
            tokens = append(tokens, token.Token{
                Kind: token.COMMA,
                Value: ",",
                Line: lineCount,
            })
        } else if isOperator(c) {
            tokens = append(tokens, token.Token{
                Kind: token.OPERATOR,
                Value: string(c),
                Line: lineCount,
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

    if l.Source[l.Cursor] == '\n' {
        lineCount++
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

func isKeyword(s string) bool {
    return slices.Index[[]string, string](keywords, s) != -1
}

func isOperator(c byte) bool {
    return c == '+' || c == '-' || c == '/' || c == '%'
}
