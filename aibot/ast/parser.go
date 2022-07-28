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
		if t.IsIdent() && t.Text == s {
			if err := ctx.advance(1); err != nil {
				return "", err
			}
			return s, nil
		}
		m, err := parseSym(ctx, s)
		if err != nil {
			return "", err
		}
		if m {
			return s, nil
		}
	}
	return "", &errorNotFound{
		msg: fmt.Sprintf("expected one of %s (%s)", strings.Join(l, ", "), t.Position),
	}
}

type errorNotFound struct {
	msg string
}

func (e *errorNotFound) Error() string { return e.msg }

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
