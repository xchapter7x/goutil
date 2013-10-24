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
	itertools.Map(s, f)
	itertools.Map(m, mf)
    fmt.Println("\n\nbegin concurrent map\n\n")
    itertools.CMap(s, f)
	itertools.CMap(m, mf)
}
