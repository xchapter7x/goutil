package itertools

import (
    "reflect"
    "fmt"
)

type Pair struct {
	First interface{}
	Second interface{}
}

func Iterate(l interface{}) (out chan Pair) {
	out = make(chan Pair)

    go func() {
        defer close(out)
        valueOfIter := reflect.ValueOf(l)

        if valueOfIter.Kind() == reflect.Ptr {
            valueOfIter = valueOfIter.Elem()
        }

        switch valueOfIter.Kind() {
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

            default:
                panic(fmt.Sprintf("Iterate function does not support the type: %s", valueOfIter.Kind()))
        }
    }()
	return
}
