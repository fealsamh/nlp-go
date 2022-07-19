package misc

import (
	"fmt"
	"reflect"
	"sort"
)

type DeepInequalityReason struct {
	IsError bool
	Text    string
	Args    []interface{}
}

func (r *DeepInequalityReason) String() string {
	if r.Args == nil {
		return r.Text
	} else {
		return fmt.Sprintf("%s (%v, %v)", r.Text, r.Args[0], r.Args[1])
	}
}

func IsDeeplyEqual(a1, a2 interface{}) (bool, *DeepInequalityReason) {
	t1 := reflect.TypeOf(a1)
	t2 := reflect.TypeOf(a2)
	if t1 != t2 {
		return false, &DeepInequalityReason{Text: "different types", Args: []interface{}{t1, t2}}
	}
	v1 := reflect.ValueOf(a1)
	v2 := reflect.ValueOf(a2)
	for t1.Kind() == reflect.Ptr {
		t1 = t1.Elem()
		t2 = t2.Elem()
		v1 = v1.Elem()
		v2 = v2.Elem()
	}
	switch t1.Kind() {
	case reflect.String:
		eq := v1.Interface().(string) == v2.Interface().(string)
		if eq {
			return true, nil
		} else {
			return false, &DeepInequalityReason{Text: "strings not equal", Args: []interface{}{v1, v2}}
		}
	case reflect.Int:
		eq := v1.Interface().(int) == v2.Interface().(int)
		if eq {
			return true, nil
		} else {
			return false, &DeepInequalityReason{Text: "ints not equal", Args: []interface{}{v1, v2}}
		}
	case reflect.Float64:
		x1 := v1.Interface().(float64)
		x2 := v2.Interface().(float64)
		eq := x1/x2 > 1-.0000001 && x1/x2 < 1+.0000001
		if eq {
			return true, nil
		} else {
			return false, &DeepInequalityReason{Text: "float64s not equal", Args: []interface{}{v1, v2}}
		}
	case reflect.Struct:
		for i := 0; i < t1.NumField(); i++ {
			a1 := v1.Field(i)
			a2 := v2.Field(i)
			eq, reason := IsDeeplyEqual(a1.Interface(), a2.Interface())
			if !eq {
				return false, reason
			}
		}
	case reflect.Slice:
		if v1.Len() != v2.Len() {
			return false, &DeepInequalityReason{Text: "different slice lengths", Args: []interface{}{v1, v2}}
		}
		for i := 0; i < v1.Len(); i++ {
			eq, reason := IsDeeplyEqual(v1.Index(i).Interface(), v2.Index(i).Interface())
			if !eq {
				return false, reason
			}
		}
	case reflect.Map:
		keys1 := GetSortedMapKeys(v1.Interface())
		keys2 := GetSortedMapKeys(v2.Interface())
		if !reflect.DeepEqual(keys1, keys2) {
			return false, &DeepInequalityReason{Text: "different map keys", Args: []interface{}{v1, v2}}
		}
		for _, k := range v1.MapKeys() {
			eq, reason := IsDeeplyEqual(v1.MapIndex(k).Interface(), v2.MapIndex(k).Interface())
			if !eq {
				return false, reason
			}
		}
	default:
		return false, &DeepInequalityReason{
			IsError: true,
			Text:    fmt.Sprintf("failed to deeply compare values of type '%v'", t1),
		}
	}
	return true, nil
}

func GetSortedMapKeys(m interface{}) interface{} {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic("expected map")
	}
	vs := v.MapKeys()
	var keys reflect.Value
	for _, v := range vs {
		if !keys.IsValid() {
			keys = reflect.MakeSlice(reflect.SliceOf(v.Type()), 0, len(vs))
		}
		keys = reflect.Append(keys, v)
	}
	switch s := keys.Interface().(type) {
	case []string:
		sort.Strings(s)
	case []int:
		sort.Ints(s)
	case []float64:
		sort.Float64s(s)
	}
	return keys.Interface()
}
