package itertools

import (
    "reflect"
    "sync"
)

func Filter( iter interface{}, f func(first, second interface{}) bool ) ( out chan Pair ){
    out = make(chan Pair)

    go func() {
        defer close(out)
        for p := range Iterate(iter) {
            args := []reflect.Value{ reflect.ValueOf(p.First), reflect.ValueOf(p.Second) }

            if reflect.ValueOf(f).Call( args )[0].Bool() {
                out <- p
            }
        }
    }()
    return
}

func CFilter( iter interface{}, f func(first, second interface{}) bool ) ( out chan Pair ){
    var wg sync.WaitGroup
    out = make(chan Pair)

    go func() {
        defer close(out)

        for p := range Iterate(iter) {
            wg.Add(1)

            go func(pp Pair) {
                defer wg.Done()
                args := []reflect.Value{ reflect.ValueOf(pp.First), reflect.ValueOf(pp.Second) }

                if reflect.ValueOf(f).Call( args )[0].Bool() {
                    out <- pp
                }
            }(p)
        }
        wg.Wait()
    }()
    return
}
