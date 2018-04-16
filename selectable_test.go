package selector_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gnampfelix/selector"
)

var _ = Describe("Selectable", func() {
	attributeFunc := func(object interface{}, key string) interface{} {
		castedObject, ok := object.(int)
		if !ok {
			return 0
		}
		switch key {
		case "square":
			return castedObject * castedObject
		case "inc":
			return castedObject - 1
		default:
			return 0
		}
	}

	It("should create a Selectable", func() {
		generator := SelectableGenerator{attributeFunc}
		selectable := generator.Generate(3)

		Expect(selectable).Should(Not(Equal(nil)))
	})

	It("should behave as a Selectable", func() {
		generator := SelectableGenerator{attributeFunc}
		selectable := generator.Generate(3)

		a1 := selectable.Attribute("square")
		a2 := selectable.Attribute("inc")
		a3 := selectable.Attribute("")
		Expect(a1).Should(Equal(9))
		Expect(a2).Should(Equal(2))
		Expect(a3).Should(Equal(0))
	})

})
