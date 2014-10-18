package unpack_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xchapter7x/goutil/unpack"
)

var _ = Describe("unpack package", func() {
	controlANew := "hi"
	controlAOld := "good"
	controlBNew := "there"
	controlBOld := "bye"

	Describe("unpack args function", func() {
		It("Should assign the value in the array to the associated pointer given", func() {
			internalA := controlAOld
			internalB := controlBOld
			arr := []interface{}{controlANew, controlBNew}
			unpack.Unpack(arr, &internalA, &internalB)
			Expect(internalA).NotTo(Equal(controlAOld))
			Expect(internalA).To(Equal(controlANew))
			Expect(internalB).NotTo(Equal(controlBOld))
			Expect(internalB).To(Equal(controlBNew))
		})
	})
})
