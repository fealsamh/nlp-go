package tokeniser

import (
	"fmt"
	"strings"
)

// TokenType ...
type TokenType int

const (
	// Word ...
	Word TokenType = iota
	// Number ...
	Number
	// Symbol ...
	Symbol
	// EOF ...
	EOF
)

func (t TokenType) String() string {
	switch t {
	case Word:
		return "WRD"
	case Number:
		return "NUM"
	case Symbol:
		return "SYM"
	case EOF:
		return "EOF"
	default:
		return "???"
	}
}

// Token ...
type Token struct {
	// The token's type.
	Type TokenType
	// The form of the token as a slice of runes.
	Form []rune
	// The line where the token is located.
	Line int
	// The column where the token is located.
	Column int
	// An associated tag.
	Tag string
}

func (t *Token) String() string {
	return fmt.Sprintf("%s/%s", string(t.Form), t.Type)
}

// Tokeniser ...
type Tokeniser struct {
	IdentChars string
}

func isWhiteChar(c rune) bool {
	return c == ' ' || c == '\r' || c == '\n' || c == '\t'
}

func (t *Tokeniser) isAlpha(c rune) bool {
	return c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= 128 || strings.ContainsRune(t.IdentChars, c)
}

func isNum(c rune) bool {
	return c >= '0' && c <= '9'
}

const (
	global = iota
	word
	number
)

// Tokenise tokenises a text.
func (t *Tokeniser) Tokenise(text string) []*Token {
	runes := []rune(text)
	var tokens []*Token
	i, line, col, colstart, state, numtag := 0, 1, 1, 1, global, ""
	var form []rune
	for {
		if state == global {
			for i < len(runes) {
				r := runes[i]
				if !isWhiteChar(r) {
					break
				}
				if r == '\n' {
					line++
					col = 1
				} else {
					col++
				}
				i++
			}
		}
		if i == len(runes) {
			break
		}
		r := runes[i]
		switch state {
		case word:
			if t.isAlpha(r) || isNum(r) {
				if numtag == "" {
					form = append(form, r)
				} else {
					numtag += string(r)
				}
				col++
				i++
			} else {
				if numtag == "" {
					tokens = append(tokens, &Token{Word, form, line, colstart, ""})
				} else {
					tokens = append(tokens, &Token{Number, form, line, colstart, numtag})
				}
				state = global
			}
		case number:
			if isNum(r) {
				form = append(form, r)
				col++
				i++
			} else {
				if t.isAlpha(r) {
					numtag += string(r)
					col++
					i++
					state = word
				} else {
					tokens = append(tokens, &Token{Number, form, line, colstart, ""})
					state = global
				}
			}
		case global:
			if t.isAlpha(r) {
				state = word
				colstart = col
				numtag = ""
				form = nil
				form = append(form, r)
				col++
				i++
			} else if isNum(r) {
				state = number
				colstart = col
				numtag = ""
				form = nil
				form = append(form, r)
				col++
				i++
			} else {
				tokens = append(tokens, &Token{Symbol, []rune{r}, line, col, ""})
				col++
				i++
			}
		}
	}
	switch state {
	case word:
		if numtag == "" {
			tokens = append(tokens, &Token{Word, form, line, colstart, ""})
		} else {
			tokens = append(tokens, &Token{Number, form, line, colstart, numtag})
		}
	case number:
		tokens = append(tokens, &Token{Number, form, line, colstart, ""})
	}
	tokens = append(tokens, &Token{EOF, nil, line, col, ""})
	return tokens
}
