package internal

import "reflect"

func IsNil[T any](val T) bool {
	rValue := reflect.ValueOf(val)
	if !rValue.IsValid() {
		return true
	}
	switch rValue.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		return rValue.IsNil()
	default:
		return false
	}
}
