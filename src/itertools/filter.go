package itertools

import (
	"reflect"
	"sync"
)

func Filter(iter interface{}, f func(first, second interface{}) bool) (out chan Pair) {
    var wg sync.WaitGroup
	out = make(chan Pair, GetIterBuffer())
    wg.Add(1)

	go func() {
		defer close(out)
        defer wg.Done()

		for p := range Iterate(iter) {
			args := []reflect.Value{reflect.ValueOf(p.First), reflect.ValueOf(p.Second)}

			if reflect.ValueOf(f).Call(args)[0].Bool() {
				out <- p
			}
		}
	}()
    wg.Wait()
	return
}

func CFilter(iter interface{}, f func(first, second interface{}) bool) (out chan Pair) {
	var wg1 sync.WaitGroup
	out = make(chan Pair, GetIterBuffer())
    wg1.Add(1)

	go func() {
		defer close(out)
        defer wg1.Done()
        var wg2 sync.WaitGroup

		for p := range Iterate(iter) {
			wg2.Add(1)

			go func(pp Pair) {
				defer wg2.Done()
				args := []reflect.Value{reflect.ValueOf(pp.First), reflect.ValueOf(pp.Second)}

				if reflect.ValueOf(f).Call(args)[0].Bool() {
					out <- pp
				}
			}(p)
		}
		wg2.Wait()
	}()
    wg1.Wait()
	return
}
