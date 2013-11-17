package itertools

import (
	"fmt"
	"reflect"
    "strings"
)

func Iterate(l interface{}) (out chan Pair) {
	out = make(chan Pair, GetIterBuffer())

	go func() {
		defer close(out)
		valueOfIter := reflect.ValueOf(l)
		k := valueOfIter.Kind()

		if k == reflect.Ptr {
			valueOfIter = valueOfIter.Elem()
			k = valueOfIter.Kind()
		}

		switch k {
		case reflect.Map:

			for _, v := range valueOfIter.MapKeys() {
				out <- Pair{v.Interface(), valueOfIter.MapIndex(v).Interface()}
			}

		case reflect.Array, reflect.Slice:

			for i := 0; i < valueOfIter.Len(); i++ {
				out <- Pair{i, valueOfIter.Index(i).Interface()}
			}

		case reflect.Chan:
			i := 0

			for v, ok := valueOfIter.Recv(); ok; {
				out <- Pair{i, v.Interface()}
				i++
			}

        case reflect.String:
            s := l.(string)

            for i, v := range strings.Split(s,"") {
                out <- Pair{i, v}
            }

		default:
			panic(fmt.Sprintf("Iterate function does not support the type: %s", k))
		}
	}()
	return
}
