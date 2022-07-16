package ast

import (
	"fmt"
	"strconv"
	"text/scanner"
)

type Expression interface {
	fmt.Stringer
	Eval(map[string]Value) (Value, error)
	Position() scanner.Position
}

type EqExpr struct {
	Lhs Expression
	Rhs Expression
	pos scanner.Position
}

func (e *EqExpr) Eval(val map[string]Value) (Value, error) {
	v1, err := e.Lhs.Eval(val)
	if err != nil {
		return nil, err
	}
	v2, err := e.Rhs.Eval(val)
	if err != nil {
		return nil, err
	}
	return &Bool{Value: v1.Equals(v2)}, nil
}

func (e *EqExpr) String() string { return fmt.Sprintf("%s = %s", e.Lhs, e.Rhs) }

func (e *EqExpr) Position() scanner.Position { return e.pos }

type IneqExpr struct {
	Lhs Expression
	Rhs Expression
	pos scanner.Position
}

func (e *IneqExpr) Eval(val map[string]Value) (Value, error) {
	v1, err := e.Lhs.Eval(val)
	if err != nil {
		return nil, err
	}
	v2, err := e.Rhs.Eval(val)
	if err != nil {
		return nil, err
	}
	return &Bool{Value: !v1.Equals(v2)}, nil
}

func (e *IneqExpr) String() string { return fmt.Sprintf("%s /= %s", e.Lhs, e.Rhs) }

func (e *IneqExpr) Position() scanner.Position { return e.pos }

type IdentExpr struct {
	Name string
	pos  scanner.Position
}

func (e *IdentExpr) Eval(val map[string]Value) (Value, error) {
	v, ok := val[e.Name]
	if !ok {
		return nil, fmt.Errorf("unknown indentifier '%s' (%s)", e.Name, e.pos)
	}
	return v, nil
}

func (e *IdentExpr) String() string { return e.Name }

func (e *IdentExpr) Position() scanner.Position { return e.pos }

type StringExpr struct {
	Value string
	pos   scanner.Position
}

func (e *StringExpr) Eval(val map[string]Value) (Value, error) {
	return &String{Value: e.Value}, nil
}

func (e *StringExpr) String() string { return strconv.Quote(e.Value) }

func (e *StringExpr) Position() scanner.Position { return e.pos }
