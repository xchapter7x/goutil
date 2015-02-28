package itertools

import (
	"errors"
	"reflect"
	"sync"
)

func CEach(iter, f interface{}) {
	if err := validateEachFunction(f); err == nil {
		var wg sync.WaitGroup

		for p := range Iterate(iter) {
			wg.Add(1)

			go func(pp Pair) {
				defer wg.Done()
				runEach(f, p)
			}(p)
		}
		wg.Wait()
	}
}

func Each(iter, f interface{}) {
	if err := validateEachFunction(f); err == nil {

		for p := range Iterate(iter) {
			runEach(f, p)
		}
	}
}

func runEach(f interface{}, p Pair) {
	function := reflect.TypeOf(f)
	pairArr := []interface{}{p.First, p.Second}
	args := []reflect.Value{}

	for i := 0; i < 2; i++ {
		switch function.NumIn() {
		case 1:
			if reflect.TypeOf(pairArr[i]).ConvertibleTo(function.In(0)) {
				arg := reflect.ValueOf(pairArr[i]).Convert(function.In(0))
				args = []reflect.Value{arg}
			}
		case 2:
			arg := reflect.ValueOf(pairArr[i]).Convert(function.In(i))
			args = append(args, arg)
		}
	}
	reflect.ValueOf(f).Call(args)
}

func validateEachFunction(f interface{}) (err error) {
	function := reflect.TypeOf(f)

	if function.Kind() != reflect.Func {
		err = errors.New("not a func type")
	}

	if function.NumIn() > 2 {
		err = errors.New("invalid argument count")
	}
	return
}
