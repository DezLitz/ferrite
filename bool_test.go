package ferrite_test

import (
	"os"

	. "github.com/dogmatiq/ferrite"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type BoolSpec", func() {
	type customBool bool

	var (
		reg  *Registry
		spec *BoolSpec[customBool]
	)

	BeforeEach(func() {
		reg = &Registry{}

		spec = BoolAs[customBool](
			"FERRITE_TEST",
			"<desc>",
			WithRegistry(reg),
		)
	})

	Describe("func Value()", func() {
		DescribeTable(
			"it returns the value associated with the literal",
			func(value string, expect customBool) {
				os.Setenv("FERRITE_TEST", value)
				defer os.Unsetenv("FERRITE_TEST")

				err := reg.Validate()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(spec.Value()).To(Equal(expect))
			},
			Entry("true", "true", customBool(true)),
			Entry("false", "false", customBool(false)),
		)
	})

	Describe("func Validate()", func() {
		When("the value is not a valid literal", func() {
			It("returns an error", func() {
				os.Setenv("FERRITE_TEST", "<invalid>")
				defer os.Unsetenv("FERRITE_TEST")

				expectErr(
					reg.Validate(),
					`ENVIRONMENT VARIABLES`,
					` ✗ FERRITE_TEST [customBool] (<desc>)`,
					`   ✓ must be set explicitly`,
					`   ✗ must be either "true" or "false", got "<invalid>"`,
				)
			})
		})

		When("the variable is not defined", func() {
			It("returns an error", func() {
				expectErr(
					reg.Validate(),
					`ENVIRONMENT VARIABLES`,
					` ✗ FERRITE_TEST [customBool] (<desc>)`,
					`   ✗ must be set explicitly`,
					`   - must be either "true" or "false"`,
				)
			})
		})
	})

	When("there is a default value", func() {
		Describe("func Value()", func() {
			When("the variable is not defined", func() {
				DescribeTable(
					"it returns the default",
					func(expect customBool) {
						spec.Default(expect)

						err := reg.Validate()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(spec.Value()).To(Equal(expect))
					},
					Entry("true", customBool(true)),
					Entry("false", customBool(false)),
				)
			})
		})
	})

	When("there are custom literals", func() {
		BeforeEach(func() {
			spec.Literals("yes", "no")
		})

		Describe("func Value()", func() {
			DescribeTable(
				"it returns the value associated with the literal",
				func(value string, expect customBool) {
					os.Setenv("FERRITE_TEST", value)
					defer os.Unsetenv("FERRITE_TEST")

					err := reg.Validate()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(spec.Value()).To(Equal(expect))
				},
				Entry("true", "yes", customBool(true)),
				Entry("false", "no", customBool(false)),
			)
		})

		Describe("func Validate()", func() {
			When("the value is not a valid literal", func() {
				It("returns an error", func() {
					os.Setenv("FERRITE_TEST", "true")
					defer os.Unsetenv("FERRITE_TEST")

					expectErr(
						reg.Validate(),
						`ENVIRONMENT VARIABLES`,
						` ✗ FERRITE_TEST [customBool] (<desc>)`,
						`   ✓ must be set explicitly`,
						`   ✗ must be either "yes" or "no", got "true"`,
					)
				})
			})
		})
	})
})
