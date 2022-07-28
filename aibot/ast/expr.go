package ast

import (
	"fmt"
	"strconv"
	"text/scanner"
)

type EvalContext struct {
	val map[string]Value
}

func (c *EvalContext) Set(k string, v Value) {
	if c.val == nil {
		c.val = make(map[string]Value)
	}
	c.val[k] = v
}

func (c *EvalContext) Get(k string) (Value, bool) {
	v, ok := c.val[k]
	return v, ok
}

type Expression interface {
	fmt.Stringer
	Eval(*EvalContext) (Value, error)
	Position() scanner.Position
}

type EqExpr struct {
	Lhs Expression
	Rhs Expression
	pos scanner.Position
}

func (e *EqExpr) Eval(ctx *EvalContext) (Value, error) {
	v1, err := e.Lhs.Eval(ctx)
	if err != nil {
		return nil, err
	}
	v2, err := e.Rhs.Eval(ctx)
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

func (e *IneqExpr) Eval(ctx *EvalContext) (Value, error) {
	v1, err := e.Lhs.Eval(ctx)
	if err != nil {
		return nil, err
	}
	v2, err := e.Rhs.Eval(ctx)
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

func (e *IdentExpr) Eval(ctx *EvalContext) (Value, error) {
	v, ok := ctx.Get(e.Name)
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

func (e *StringExpr) Eval(ctx *EvalContext) (Value, error) {
	return &String{Value: e.Value}, nil
}

func (e *StringExpr) String() string { return strconv.Quote(e.Value) }

func (e *StringExpr) Position() scanner.Position { return e.pos }

type IntExpr struct {
	Value int
	pos   scanner.Position
}

func (e *IntExpr) Eval(ctx *EvalContext) (Value, error) {
	return &Int{Value: e.Value}, nil
}

func (e *IntExpr) String() string { return strconv.Itoa(e.Value) }

func (e *IntExpr) Position() scanner.Position { return e.pos }

type FloatExpr struct {
	Value float64
	pos   scanner.Position
}

func (e *FloatExpr) Eval(ctx *EvalContext) (Value, error) {
	return &Float{Value: e.Value}, nil
}

func (e *FloatExpr) String() string { return fmt.Sprintf("%f", e.Value) }

func (e *FloatExpr) Position() scanner.Position { return e.pos }
