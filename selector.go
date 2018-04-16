package selector

// A Selector is a collection of criteria that a Selectable can be
// evaluated against.
type Selector interface {
	// Match checks if the Selectable meets the criteria the Selector bundles.
	// Each implementation should also handle the case where the Selectable is
	// nil.
	Match(to Selectable) bool
}

type anySelector bool

// Any returns a selector that will match to any Selectable
func Any() Selector {
	return anySelector(true)
}

func (a anySelector) Match(to Selectable) bool {
	return true
}

type noneSelector bool

// None returns a selector that won't match to any Selectable
func None() Selector {
	return noneSelector(true)
}

func (n noneSelector) Match(to Selectable) bool {
	return false
}

type allSelector struct {
	selector Selector
}

// All returns a selector that will only match a selectable, if the selectable is
// an array and all of its members match the inner selector.
func All(selector Selector) Selector {
	return allSelector{selector}
}

func (a allSelector) Match(to Selectable) bool {
	if to == nil {
		return true
	}

	selectables, ok := to.(Selectables)
	if !ok {
		return false
	}

	for _, selectable := range selectables {
		if !a.selector.Match(selectable) {
			return false
		}
	}
	return true
}

type firstSelector struct {
	selector Selector
}

// First returns a selector that will match an array of selecatbles, if at least
// one element of the array matches the inner selector.
func First(selector Selector) Selector {
	return firstSelector{selector}
}

func (f firstSelector) Match(to Selectable) bool {
	if to == nil {
		return false
	}

	selectables, ok := to.(Selectables)
	if !ok {
		return false
	}

	for _, selectable := range selectables {
		if f.selector.Match(selectable) {
			return true
		}
	}
	return false
}

type inverseSelector struct {
	selector Selector
}

// Inverse returns a Selector that will match a Selectable if the inner Selector
// won't
func Inverse(selector Selector) Selector {
	return inverseSelector{selector}
}

func (i inverseSelector) Match(to Selectable) bool {
	return !i.selector.Match(to)
}

type lambdaSelector struct {
	match func(to Selectable) bool
}

// Lamdba returns a Selector that will match when the given function returns true.
func Lambda(match func(to Selectable) bool) Selector {
	return lambdaSelector{match}
}

func (l lambdaSelector) Match(to Selectable) bool {
	return l.match(to)
}
