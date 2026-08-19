package ast

import (
	"fmt"
	"strconv"
	"text/scanner"
)

// EvalContext ...
type EvalContext struct {
	val map[string]Value
}

// Set ...
func (c *EvalContext) Set(k string, v Value) {
	if c.val == nil {
		c.val = make(map[string]Value)
	}
	c.val[k] = v
}

// Get ...
func (c *EvalContext) Get(k string) (Value, bool) {
	v, ok := c.val[k]
	return v, ok
}

// Expression ...
type Expression interface {
	fmt.Stringer
	Eval(*EvalContext) (Value, error)
	Position() scanner.Position
}

// EqExpr ...
type EqExpr struct {
	LHS Expression
	RHS Expression
	pos scanner.Position
}

// Eval ...
func (e *EqExpr) Eval(ctx *EvalContext) (Value, error) {
	v1, err := e.LHS.Eval(ctx)
	if err != nil {
		return nil, err
	}
	v2, err := e.RHS.Eval(ctx)
	if err != nil {
		return nil, err
	}
	return &Bool{Value: v1.Equals(v2)}, nil
}

func (e *EqExpr) String() string { return fmt.Sprintf("%s = %s", e.LHS, e.RHS) }

// Position ...
func (e *EqExpr) Position() scanner.Position { return e.pos }

// IneqExpr ...
type IneqExpr struct {
	LHS Expression
	RHS Expression
	pos scanner.Position
}

// Eval ...
func (e *IneqExpr) Eval(ctx *EvalContext) (Value, error) {
	v1, err := e.LHS.Eval(ctx)
	if err != nil {
		return nil, err
	}
	v2, err := e.RHS.Eval(ctx)
	if err != nil {
		return nil, err
	}
	return &Bool{Value: !v1.Equals(v2)}, nil
}

func (e *IneqExpr) String() string { return fmt.Sprintf("%s /= %s", e.LHS, e.RHS) }

// Position ...
func (e *IneqExpr) Position() scanner.Position { return e.pos }

// IdentExpr ...
type IdentExpr struct {
	Name string
	pos  scanner.Position
}

// Eval ...
func (e *IdentExpr) Eval(ctx *EvalContext) (Value, error) {
	v, ok := ctx.Get(e.Name)
	if !ok {
		return nil, fmt.Errorf("unknown indentifier '%s' (%s)", e.Name, e.pos)
	}
	return v, nil
}

func (e *IdentExpr) String() string { return e.Name }

// Position ...
func (e *IdentExpr) Position() scanner.Position { return e.pos }

// StringExpr ...
type StringExpr struct {
	Value string
	pos   scanner.Position
}

// Eval ...
func (e *StringExpr) Eval(ctx *EvalContext) (Value, error) {
	return &String{Value: e.Value}, nil
}

func (e *StringExpr) String() string { return strconv.Quote(e.Value) }

// Position ...
func (e *StringExpr) Position() scanner.Position { return e.pos }

// IntExpr ...
type IntExpr struct {
	Value int
	pos   scanner.Position
}

// Eval ...
func (e *IntExpr) Eval(ctx *EvalContext) (Value, error) {
	return &Int{Value: e.Value}, nil
}

// String ...
func (e *IntExpr) String() string { return strconv.Itoa(e.Value) }

// Position ...
func (e *IntExpr) Position() scanner.Position { return e.pos }

// FloatExpr ...
type FloatExpr struct {
	Value float64
	pos   scanner.Position
}

// Eval ...
func (e *FloatExpr) Eval(ctx *EvalContext) (Value, error) {
	return &Float{Value: e.Value}, nil
}

// String ...
func (e *FloatExpr) String() string { return fmt.Sprintf("%f", e.Value) }

// Position ...
func (e *FloatExpr) Position() scanner.Position { return e.pos }
