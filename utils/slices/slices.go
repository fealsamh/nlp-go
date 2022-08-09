package slices

import "reflect"

func Contains(slice, el interface{}) bool {
	slcval := reflect.ValueOf(slice)
	if slcval.Kind() != reflect.Slice {
		panic("slices.Contains: first argument not a slice")
	}

	elval := reflect.ValueOf(el)
	if slcval.Type().Elem() != elval.Type() {
		panic("slicesContains: first argument's element type not matching second argument's type")
	}

	for i := 0; i < slcval.Len(); i++ {
		if reflect.DeepEqual(slcval.Index(i).Interface(), elval.Interface()) {
			return true
		}
	}

	return false
}
