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
	Interface() any
	Equals(Value) bool
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
func (v *String) Interface() any { return v.Value }

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
func (v *Bool) Interface() any { return v.Value }

// Equals ...
func (v *Bool) Equals(v2 Value) bool {
	if v2, ok := v2.(*Bool); ok {
		return v.Value == v2.Value
	}
	return false
}
