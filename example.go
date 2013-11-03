package main

import (
    "fmt"
    "itertools"
)

func f(i int, v string) string {
	fmt.Println( i, v )
	return v
}

func mf(i, v string) string {
	fmt.Println( i, v )
	return v
}

func main() {
    s := []string{"asdf","asdfasdf","geeeg","gggggggg"}
    m := map[string]string{"a":"asdf","b":"asdfasdf","c":"geeeg","d":"gggggggg"}
	itertools.Map(&s, f)
	itertools.Map(&m, mf)
    fmt.Println("\n\nbegin concurrent map\n\n")
    itertools.CMap(s, f)
	itertools.CMap(m, mf)

    fmt.Println("\n\nFilter Sample\n\n")


    f := itertools.Filter(s, func(i , v interface{}) bool {
        il := map[int]int{1:1,2:2}
        _, ok := il[i.(int)]
        return ok
    })

    for i := range f {
        fmt.Println(i)
    }

    fmt.Println("\n\nConcurrent Filter Sample\n\n")

    fC := itertools.CFilter(s, func(i , v interface{}) bool {
        il := map[int]int{1:1,2:2}
        _, ok := il[i.(int)]
        return ok
    })

    for i := range fC {
        fmt.Println(i)
    }
}
