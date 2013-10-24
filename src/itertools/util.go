package itertools

import (
    "reflect"
    "sync"
)

func CMap( iter, f interface{} ) {
    var wg sync.WaitGroup
    iterV := reflect.ValueOf(iter)

    switch iterV.Kind() {
        case reflect.Map:
            valueOfIter := reflect.ValueOf(iter)

            for _, v := range valueOfIter.MapKeys() {
                wg.Add(1)
                go func(vv reflect.Value) {
                    defer wg.Done()
                    reflect.ValueOf(f).Call( []reflect.Value{ vv,  valueOfIter.MapIndex(vv) } )
                }(v)
            }

        case reflect.Array, reflect.Slice:
            valueOfIter := reflect.ValueOf(iter)

            for i := 0; i < valueOfIter.Len(); i++ {
                wg.Add(1)
                go func(ii int) {
                    defer wg.Done()
                    reflect.ValueOf(f).Call( []reflect.Value{ reflect.ValueOf(ii),  valueOfIter.Index(ii) } )
                }(i)
            }
    }
    wg.Wait()
}

func Map( iter, f interface{} ) {
    iterV := reflect.ValueOf(iter)
    switch iterV.Kind() {
        case reflect.Map:
            valueOfIter := reflect.ValueOf(iter)

            for _, v := range valueOfIter.MapKeys() {
                reflect.ValueOf(f).Call( []reflect.Value{ v,  valueOfIter.MapIndex(v) } )
            }

        case reflect.Array, reflect.Slice:
            valueOfIter := reflect.ValueOf(iter)

            for i := 0; i < valueOfIter.Len(); i++ {
                reflect.ValueOf(f).Call( []reflect.Value{ reflect.ValueOf(i),  valueOfIter.Index(i) } )
            }
    }
}
