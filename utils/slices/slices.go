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

func Equal(slice1, slice2 interface{}) bool {
	slcval1 := reflect.ValueOf(slice1)
	if slcval1.Kind() != reflect.Slice {
		panic("slices.Equal: first argument not a slice")
	}

	slcval2 := reflect.ValueOf(slice2)
	if slcval2.Kind() != reflect.Slice {
		panic("slices.Equal: second argument not a slice")
	}

	if slcval1.Type().Elem() != slcval2.Type().Elem() {
		panic("slices.Equal: arguments' element types do not match")
	}

	if slcval1.Len() != slcval2.Len() {
		return false
	}

	for i := 0; i < slcval1.Len(); i++ {
		if !reflect.DeepEqual(slcval1.Index(i).Interface(), slcval2.Index(i).Interface()) {
			return false
		}
	}

	return true
}
