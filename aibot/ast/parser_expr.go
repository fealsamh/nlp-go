package ast

import (
	"fmt"
)

// ParseExpr ...
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
