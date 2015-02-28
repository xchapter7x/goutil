package itertools

import (
	"errors"
	"reflect"
	"sync"
)

var (
	NotFuncError    = errors.New("not a func type")
	InvalidArgError = errors.New("invalid argument count")
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
	maxArgLen := 2
	function := reflect.TypeOf(f)
	pairArr := []interface{}{p.First, p.Second}
	args := []reflect.Value{}

	for i := 0; i < maxArgLen; i++ {
		switch function.NumIn() {
		case 1:
			firstArg := 0

			if reflect.TypeOf(pairArr[i]).ConvertibleTo(function.In(firstArg)) {
				arg := reflect.ValueOf(pairArr[i]).Convert(function.In(firstArg))
				args = []reflect.Value{arg}
			}
		case maxArgLen:
			arg := reflect.ValueOf(pairArr[i]).Convert(function.In(i))
			args = append(args, arg)
		}
	}
	reflect.ValueOf(f).Call(args)
}

func validateEachFunction(f interface{}) (err error) {
	function := reflect.TypeOf(f)

	if function.Kind() != reflect.Func {
		err = NotFuncError
	}

	if function.NumIn() > 2 {
		err = InvalidArgError
	}
	return
}
