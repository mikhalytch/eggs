package deref

// OrDefault is a 'dereference or default':
// it will dereference the pointer passed, or return default value for type T if p == nil.
func OrDefault[T any](ptr *T) T {
	if ptr == nil {
		ptr := new(T)

		return *ptr
	}

	return *ptr
}

// Of is an alias to OrDefault.
func Of[T any](p *T) T { return OrDefault(p) }
