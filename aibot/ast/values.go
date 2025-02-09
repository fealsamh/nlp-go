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

// Must ...
func Must(v Value, err error) Value {
	if err != nil {
		panic(err)
	}
	return v
}

// NewValue ...
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

// Int ...
type Int struct {
	Value int
}

// Kind ...
func (v *Int) Kind() reflect.Kind { return reflect.Int }

// String ...
func (v *Int) String() string { return strconv.Itoa(v.Value) }

// Interface ...
func (v *Int) Interface() interface{} { return v.Value }

// Equals ...
func (v *Int) Equals(v2 Value) bool {
	if v2, ok := v2.(*Int); ok {
		return v.Value == v2.Value
	}
	return false
}

// Float ...
type Float struct {
	Value float64
}

// Kind ...
func (v *Float) Kind() reflect.Kind { return reflect.Float64 }

// String ...
func (v *Float) String() string { return fmt.Sprintf("%f", v.Value) }

// Interface ...
func (v *Float) Interface() interface{} { return v.Value }

// Equals ...
func (v *Float) Equals(v2 Value) bool {
	if v2, ok := v2.(*Float); ok {
		return v.Value == v2.Value
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
