package itertools

import (
    "testing"
)

var f_called int = 0
var mf_called int = 0

func f(i int, v string) string {
	f_called += 1
	return v
}

func mf(i, v string) string {
	mf_called += 1
	return v
}

func Test_MapSliceArray(t *testing.T) {
    s := []string{"asdf","asdfasdf","geeeg","gggggggg"}
    Map(s, f)

    if f_called != len(s) {
        t.Errorf("func f was not called %d times", len(s))
    }

    f_called = 0
}

func Test_MapMap(t *testing.T) {
    m := map[string]string{"a":"asdf","b":"asdfasdf","c":"geeeg","d":"gggggggg"}
	Map(m, mf)

    if mf_called != len(m) {
        t.Errorf("func mf was not called %d times", len(m))
    }
    mf_called = 0
}

func Test_CMapSliceArray(t *testing.T) {
    s := []string{"asdf","asdfasdf","geeeg","gggggggg"}
    CMap(s, f)

    if f_called != len(s) {
        t.Errorf("func f was not called %d times", len(s))
    }
    f_called = 0
}

func Test_CMapMap(t *testing.T) {
    m := map[string]string{"a":"asdf","b":"asdfasdf","c":"geeeg","d":"gggggggg"}
	CMap(m, mf)

    if mf_called != len(m) {
        t.Errorf("func mf was not called %d times", len(m))
    }
    mf_called = 0
}
