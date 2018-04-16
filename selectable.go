package selector

// Selectable provides an interface for the Selector to retrieve an object's criteria
// without using reflection.
type Selectable interface {
	// Attribute returns the named attribute of the Selectable
	Attribute(key string) interface{}
}

// Selectables are needed by some of the Selectors provided in this package. They
// make arrays of type Selectable implement the interface as well.
type Selectables []Selectable

// Selectables implement the Selectable interface.
func (s Selectables) Attribute(key string) interface{} {
	return nil
}

type selectableInterface struct {
	payload       interface{}
	attributeFunc func(interface{}, string) interface{}
}

func (s selectableInterface) Attribute(key string) interface{} {
	return s.attributeFunc(s.payload, key)
}

// A SelectableGenerator enables the user to convert any object into a Selectable.
// An instance must have an AttributeFunc which will be evaluated whenever the
// the Selectable's Attribute Function is called.
type SelectableGenerator struct {
	AttributeFunc func(interface{}, string) interface{}
}

// Generate creates a new Selectable by chaining the payload to the AttributeFunc
// of the Generator.
func (s SelectableGenerator) Generate(payload interface{}) Selectable {
	return selectableInterface{
		payload:       payload,
		attributeFunc: s.AttributeFunc,
	}
}
