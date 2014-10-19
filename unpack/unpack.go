package unpack

import (
	"fmt"
	"reflect"
)

func Unpack(arr []interface{}, args ...interface{}) (err error) {
	return UnpackArray(arr, args)
}

func UnpackArray(arr []interface{}, args []interface{}) (err error) {

	if len(arr) == len(args) {

		for i, v := range arr {
			ptrVal := reflect.ValueOf((args)[i])
			ptrElem := ptrVal.Elem()

			if ptrElem.Kind() == reflect.ValueOf(v).Kind() {
				ptrElem.Set(reflect.ValueOf(v))
				(args)[i] = ptrElem.Interface()

			} else {
				err = fmt.Errorf("Incorrect pointer type %s != %s", ptrElem.Kind(), reflect.ValueOf(v).Kind())
			}
		}

	} else {
		err = fmt.Errorf("Incorrect argument count: pointers dont match response element count %s != %s", len(arr), len(args))
	}
	return
}
