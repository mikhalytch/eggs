package deref

import (
	"github.com/mikhalytch/eggs/dflt"
)

// OrDefault is a 'dereference or default':
// it will dereference the pointer passed, or return default value for type T if p == nil.
func OrDefault[T any](ptr *T) T {
	if ptr == nil {
		return dflt.Of[T]()
	}

	return *ptr
}

// Of is an alias to OrDefault.
func Of[T any](p *T) T { return OrDefault(p) }
