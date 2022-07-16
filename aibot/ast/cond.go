package ast

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type parsingCtx struct {
	tokens []*Token
}

func (c *parsingCtx) token(i int) *Token { return c.tokens[i] }

func (c *parsingCtx) advance(n int) error {
	if n == 0 {
		return nil
	}
	if c.token(0).IsEOF() {
		return errors.New("unexpected end of file")
	}
	c.tokens = c.tokens[1:]
	return c.advance(n - 1)
}

func ParseExpr(tokens []*Token) (Expression, error) {
	return parseComp(&parsingCtx{
		tokens: tokens,
	})
}

func parseComp(ctx *parsingCtx) (Expression, error) {
	pos := ctx.token(0).Position

	e1, err := parseAtom(ctx)
	if err != nil {
		return nil, err
	}

	s, err := parseOneOf(ctx, "=", "/=")
	if err != nil {
		return nil, err
	}

	e2, err := parseAtom(ctx)
	if err != nil {
		return nil, err
	}

	switch s {
	case "=":
		return &EqExpr{Lhs: e1, Rhs: e2, pos: pos}, nil
	case "/=":
		return &IneqExpr{Lhs: e1, Rhs: e2, pos: pos}, nil
	}

	panic("arrived at unreachable branch")
}

func parseAtom(ctx *parsingCtx) (Expression, error) {
	t := ctx.token(0)
	switch {
	case t.IsIdent():
		if err := ctx.advance(1); err != nil {
			return nil, err
		}
		return &IdentExpr{Name: t.Text, pos: t.Position}, nil
	case t.IsString():
		if err := ctx.advance(1); err != nil {
			return nil, err
		}
		return &StringExpr{Value: t.Value().(string), pos: t.Position}, nil
	default:
		return nil, fmt.Errorf("unexpected token (%s)", t.Position)
	}
}

func parseOneOf(ctx *parsingCtx, l ...string) (string, error) {
	sort.Slice(l, func(i, j int) bool {
		s1, s2 := l[i], l[j]
		if len(s1) > len(s2) {
			return true
		}
		if len(s1) < len(s2) {
			return false
		}
		return s1 < s2
	})
	t := ctx.token(0)
	for _, s := range l {
		m, err := parseSym(ctx, s)
		if err != nil {
			return "", err
		}
		if m {
			return s, nil
		}
	}
	return "", fmt.Errorf("expected one of %s (%s)", strings.Join(l, ", "), t.Position)
}

func parseSym(ctx *parsingCtx, s string) (bool, error) {
	for i := 0; i < len(s); i++ {
		t := ctx.token(i)
		if t.IsEOF() {
			return false, errors.New("unexpected end of file")
		}
		if !t.IsSymbol() || t.Text != s[i:i+1] {
			return false, nil
		}
	}
	if err := ctx.advance(len(s)); err != nil {
		return false, err
	}
	return true, nil
}
