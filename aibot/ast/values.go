package ast

import (
	"fmt"
	"reflect"
	"strconv"
)

// Value ...
type Value interface {
	fmt.Stringer
	Kind() reflect.Kind
	Interface() interface{}
	Equals(Value) bool
}

func Must(v Value, err error) Value {
	if err != nil {
		panic(err)
	}
	return v
}

func NewValue(x interface{}) (Value, error) {
	switch x := x.(type) {
	case bool:
		return &Bool{Value: x}, nil
	case string:
		return &String{Value: x}, nil
	case int:
		return &Int{Value: x}, nil
	case float64:
		return &Float{Value: x}, nil
	}
	return nil, fmt.Errorf("failed to created a value of type '%T'", x)
}

type Int struct {
	Value int
}

func (v *Int) Kind() reflect.Kind { return reflect.Int }

func (v *Int) String() string { return strconv.Itoa(v.Value) }

func (v *Int) Interface() interface{} { return v.Value }

func (v1 *Int) Equals(v2 Value) bool {
	if v2, ok := v2.(*Int); ok {
		return v1.Value == v2.Value
	}
	return false
}

type Float struct {
	Value float64
}

func (v *Float) Kind() reflect.Kind { return reflect.Float64 }

func (v *Float) String() string { return fmt.Sprintf("%f", v.Value) }

func (v *Float) Interface() interface{} { return v.Value }

func (v1 *Float) Equals(v2 Value) bool {
	if v2, ok := v2.(*Float); ok {
		return v1.Value == v2.Value
	}
	return false
}

// String ...
type String struct {
	Value string
}

// Kind ...
func (v *String) Kind() reflect.Kind { return reflect.String }

// String ...
func (v *String) String() string { return strconv.Quote(v.Value) }

// Interface ...
func (v *String) Interface() interface{} { return v.Value }

// Equals ...
func (v *String) Equals(v2 Value) bool {
	if v2, ok := v2.(*String); ok {
		return v.Value == v2.Value
	}
	return false
}

// Bool ...
type Bool struct {
	Value bool
}

// Kind ...
func (v *Bool) Kind() reflect.Kind { return reflect.Bool }

// String ...
func (v *Bool) String() string {
	if v.Value {
		return "true"
	}
	return "false"
}

// Interface ...
func (v *Bool) Interface() interface{} { return v.Value }

// Equals ...
func (v *Bool) Equals(v2 Value) bool {
	if v2, ok := v2.(*Bool); ok {
		return v.Value == v2.Value
	}
	return false
}
