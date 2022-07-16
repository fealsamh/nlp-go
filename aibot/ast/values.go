package ast

import (
	"reflect"
)

type Value interface {
	Kind() reflect.Kind
	Interface() interface{}
	Equals(Value) bool
}

type String struct {
	Value string
}

func (v *String) Kind() reflect.Kind { return reflect.String }

func (v *String) Interface() interface{} { return v.Value }

func (v1 *String) Equals(v2 Value) bool {
	if v2, ok := v2.(*String); ok {
		return v1.Value == v2.Value
	}
	return false
}

type Bool struct {
	Value bool
}

func (v *Bool) Kind() reflect.Kind { return reflect.Bool }

func (v *Bool) Interface() interface{} { return v.Value }

func (v1 *Bool) Equals(v2 Value) bool {
	if v2, ok := v2.(*Bool); ok {
		return v1.Value == v2.Value
	}
	return false
}
