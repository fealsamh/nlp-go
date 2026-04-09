package ast

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// Token ...
type Token struct {
	tok      rune
	Text     string
	Position scanner.Position
}

// Kind ...
func (t *Token) Kind() string { return scanner.TokenString(t.tok) }

// IsSymbol ...
func (t *Token) IsSymbol() bool { return scanner.TokenString(t.tok)[0] == '"' }

// Value ...
func (t *Token) Value() any {
	switch t.tok {
	case scanner.String, scanner.RawString:
		s, err := strconv.Unquote(t.Text)
		if err != nil {
			panic(err)
		}
		return s
	case scanner.Int:
		x, err := strconv.Atoi(t.Text)
		if err != nil {
			panic(err)
		}
		return x
	case scanner.Float:
		x, err := strconv.ParseFloat(t.Text, 64)
		if err != nil {
			panic(err)
		}
		return x
	default:
		return t.Text
	}
}

// IsIdent ...
func (t *Token) IsIdent() bool { return t.tok == scanner.Ident }

// IsString ...
func (t *Token) IsString() bool { return t.tok == scanner.String || t.tok == scanner.RawString }

// IsInt ...
func (t *Token) IsInt() bool { return t.tok == scanner.Int }

// IsFloat ...
func (t *Token) IsFloat() bool { return t.tok == scanner.Float }

// IsEOF ...
func (t *Token) IsEOF() bool { return t.tok == scanner.EOF }

// Tokenise ...
func Tokenise(code string) ([]*Token, error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(code))
	var err error
	s.Error = func(s *scanner.Scanner, msg string) {
		err = fmt.Errorf("%s (%s)", msg, s.Position)
	}
	var tokens []*Token
	for {
		tok := s.Scan()
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, &Token{
			tok:      tok,
			Text:     s.TokenText(),
			Position: s.Position,
		})
		if tok == scanner.EOF {
			break
		}
	}
	return tokens, nil
}
