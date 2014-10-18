package unpack

import (
	"reflect"
)

func Unpack(arr []interface{}, args ...interface{}) {
	UnpackArray(arr, args)
}

func UnpackArray(arr []interface{}, args []interface{}) {
	for i, v := range arr {
		ptrVal := reflect.ValueOf((args)[i])
		ptrElem := ptrVal.Elem()
		ptrElem.Set(reflect.ValueOf(v))
		(args)[i] = ptrElem.Interface()
	}
}
