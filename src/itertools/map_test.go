package itertools

import (
	"testing"
)

var (
	f_called int
	s        []string
	m        map[string]string
)

func Setup() {
	f_called = 0
	s = []string{"asdf", "asdfasdf", "geeeg", "gggggggg"}
	m = map[string]string{"a": "asdf", "b": "asdfasdf", "c": "geeeg", "d": "gggggggg"}
}

func TearDown() {
	Setup()
}

func Test_MapSliceArray(t *testing.T) {
	Setup()
	defer TearDown()

	Map(s, func(i int, v string) string {
		f_called += 1
		return v
	})

	if f_called != len(s) {
		t.Errorf("func f was not called %d times", len(s))
	}
}

func Test_MapMap(t *testing.T) {
	Setup()
	defer TearDown()

	Map(m, func(i, v string) string {
		f_called += 1
		return v
	})

	if f_called != len(m) {
		t.Errorf("func mf was not called %d times", len(m))
	}
}

func Test_CMapSliceArray(t *testing.T) {
	Setup()
	defer TearDown()

	CMap(s, func(i int, v string) string {
		f_called += 1
		return v
	})

	if f_called != len(s) {
		t.Errorf("func f was not called %d times", len(s))
	}
}

func Test_CMapMap(t *testing.T) {
	Setup()
	defer TearDown()

	CMap(m, func(i, v string) string {
		f_called += 1
		return v
	})

	if f_called != len(m) {
		t.Errorf("func mf was not called %d times", len(m))
	}
}
