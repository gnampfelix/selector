package selector_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gnampfelix/selector"
)

var _ = Describe("Selector", func() {
	valuefunc := func(object interface{}, key string) interface{} {
		return object
	}

	generator := SelectableGenerator{valuefunc}

	s1 := generator.Generate("a")
	s2 := generator.Generate("b")
	s3 := generator.Generate("c")

	selectables := Selectables{s1, s2, s3}

	lambda := func(to Selectable) bool {
		if to == nil {
			return false
		}

		attribute := to.Attribute("")
		converted, ok := attribute.(string)

		if !ok {
			return false
		}

		return converted == "a"
	}

	It("should match lambda", func() {
		selector := Lambda(lambda)
		b1 := selector.Match(s1)
		b2 := selector.Match(s2)

		Expect(b1).Should(BeTrue())
		Expect(b2).Should(BeFalse())
	})

	It("should match not", func() {
		selector := Inverse(Lambda(lambda))

		b3 := selector.Match(s3)
		Expect(b3).Should(BeTrue())
	})

	It("should match first", func() {
		selector := First(Lambda(lambda))

		b4 := selector.Match(selectables)
		Expect(b4).Should(BeTrue())
	})

	It("should not match all", func() {
		selector := All(Lambda(lambda))

		b5 := selector.Match(selectables)
		Expect(b5).Should(BeFalse())
	})

	It("should match none", func() {
		selector := All(None())
		b6 := selector.Match(selectables)
		Expect(b6).Should(BeFalse())
	})

	It("should match any", func() {
		selector := Any()
		b7 := selector.Match(selectables)
		Expect(b7).Should(BeTrue())
	})

})
