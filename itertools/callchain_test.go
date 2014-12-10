package itertools_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/goutil/itertools"
)

var _ = Describe("CallChain", func() {
	var callCount int = 0
	var sampleSuccessReturn string = "success"
	var sampleFailureReturn string = "failure"
	var controlError error = fmt.Errorf(sampleFailureReturn)
	BeforeEach(func() {
		callCount = 0
	})

	AfterEach(func() {
		callCount = 0
	})

	var successMultiReturn func(string) (string, error) = func(s string) (string, error) {
		callCount++
		return sampleSuccessReturn, nil
	}

	var failMultiReturn func(string) (string, error) = func(s string) (string, error) {
		callCount++
		return sampleFailureReturn, controlError
	}

	var successNoArgMultiReturn func() (string, error) = func() (string, error) {
		return sampleSuccessReturn, nil
	}

	var failNoArgMultiReturn func() (string, error) = func() (string, error) {
		return sampleFailureReturn, controlError
	}

	var runNoError func() string = func() string {
		return sampleSuccessReturn
	}

	Context("CallChain function", func() {
		Context("with a non nil chained error", func() {
			It("Should return a error equal to the chained error", func() {
				e := fmt.Errorf("new error")
				_, err := CallChain(e, failNoArgMultiReturn, "testing_error")
				Ω(err).ShouldNot(BeNil())
				Ω(err).Should(Equal(e))
			})

			It("Should return an error, skipping any call including failed calls, and return the original error", func() {
				e := fmt.Errorf("new error")
				_, err := CallChain(e, successMultiReturn, "testing_error")
				Ω(err).ShouldNot(Equal(controlError))
			})

			It("Should skip the function if passed an error - even a success call function", func() {
				e := fmt.Errorf("new error")
				CallChain(e, successMultiReturn, "testing_error")
				Ω(callCount).Should(Equal(0))
			})

			It("Should skip the function if passed an error - even a failed call function", func() {
				e := fmt.Errorf("new error")
				CallChain(e, successMultiReturn, "testing_error")
				Ω(callCount).Should(Equal(0))
			})

		})

		Context("with a nil chained error", func() {
			Context("on success", func() {
				It("Should return a nil error w/ multiple return functions and arguments", func() {
					_, err := CallChain(nil, successMultiReturn, "testing_error")
					Ω(err).Should(BeNil())
				})

				It("Should return a nil error w/ no arguments", func() {
					_, err := CallChain(nil, successNoArgMultiReturn)
					Ω(err).Should(BeNil())
				})

				It("Should return a nil error w/ error values returned", func() {
					_, err := CallChain(nil, runNoError)
					Ω(err).Should(BeNil())
				})
			})

			Context("on failure", func() {
				It("Should return a non nil error w/ multiple return values", func() {
					_, err := CallChain(nil, failMultiReturn, "testing_error")
					Ω(err).ShouldNot(BeNil())
					Ω(err).Should(Equal(controlError))
				})

				It("Should return a non nil error w/ no args", func() {
					_, err := CallChain(nil, failNoArgMultiReturn)
					Ω(err).ShouldNot(BeNil())
					Ω(err).Should(Equal(controlError))
				})
			})
		})
	})
})
