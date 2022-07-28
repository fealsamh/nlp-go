package ast

import (
	"fmt"
)

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
		if err, ok := err.(ParseError); ok && (err.Kind() == errorKindNotFound || err.Kind() == errorKindEOF) {
			return e1, nil
		}
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
	case t.IsInt():
		if err := ctx.advance(1); err != nil {
			return nil, err
		}
		return &IntExpr{Value: t.Value().(int), pos: t.Position}, nil
	case t.IsFloat():
		if err := ctx.advance(1); err != nil {
			return nil, err
		}
		return &FloatExpr{Value: t.Value().(float64), pos: t.Position}, nil
	default:
		return nil, fmt.Errorf("unexpected token (%s)", t.Position)
	}
}
