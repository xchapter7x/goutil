package itertools_test

import (
	"strings"
	"testing"

	. "github.com/xchapter7x/goutil/itertools"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filter", func() {
	var iterable []string = []string{
		"hi",
		"there",
		"who are",
		"you",
		"that thing",
	}
	var output chan Pair

	Context("called w/ non standard function format", func() {
		BeforeEach(func() {
			output = Filter(iterable, func(i int, v string) bool {
				return strings.HasPrefix(v, "t")
			})
		})

		It("should contain the correct number of elements", func() {
			Ω(len(output)).Should(Equal(2))
		})

		It("should contain a value matches from the control iterable", func() {
			val1 := <-output
			val2 := <-output
			Ω(arrayContains(val1.Second.(string), iterable)).Should(BeTrue())
			Ω(arrayContains(val2.Second.(string), iterable)).Should(BeTrue())
		})
	})
})

func arrayContains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

var filterTestData []string
var filterWhiteList []string

func SetupFilter() {
	filterTestData = []string{"asdf", "asdfasdf", "geeeg", "gggggggg"}
	filterWhiteList = []string{"asdfasdf", "geeeg"}
}

func TearDownFilter() {
	filterTestData = []string{}
	filterWhiteList = []string{}
}

func Test_CFilter(t *testing.T) {
	SetupFilter()
	defer TearDownFilter()

	f := CFilter(filterTestData, func(i, v interface{}) bool {
		return findInStringArray(v.(string), filterWhiteList)
	})

	for r := range f {

		if !findInStringArray(r.Second.(string), filterWhiteList) {
			t.Errorf("Error: %s should have been filtered, but it was not ", r)
		}
	}
}

func Test_Filter(t *testing.T) {
	SetupFilter()
	defer TearDownFilter()

	f := Filter(filterTestData, func(i, v interface{}) bool {
		return findInStringArray(v.(string), filterWhiteList)
	})

	for r := range f {

		if !findInStringArray(r.Second.(string), filterWhiteList) {
			t.Errorf("Error: %s should have been filtered, but it was not ", r)
		}
	}
}

func Test_CFilterFalse(t *testing.T) {
	SetupFilter()
	defer TearDownFilter()

	f := CFilterFalse(filterTestData, func(i, v interface{}) bool {
		return findInStringArray(v.(string), filterWhiteList)
	})

	for r := range f {

		if findInStringArray(r.Second.(string), filterWhiteList) {
			t.Errorf("Error: %s should have been filtered, but it was not ", r)
		}
	}
}

func Test_FilterFalse(t *testing.T) {
	SetupFilter()
	defer TearDownFilter()

	f := FilterFalse(filterTestData, func(i, v interface{}) bool {
		return findInStringArray(v.(string), filterWhiteList)
	})

	for r := range f {

		if findInStringArray(r.Second.(string), filterWhiteList) {
			t.Errorf("Error: %s should have been filtered, but it was not ", r)
		}
	}
}

func findInStringArray(v string, a []string) (r bool) {
	r = false

	for _, i := range a {

		if v == i {
			r = true
		}
	}
	return
}
