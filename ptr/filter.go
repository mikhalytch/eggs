package ptr

import (
	"github.com/mikhalytch/eggs/funcs"
	"github.com/mikhalytch/eggs/funcs/predicate"
)

func Filter[T any](tPtr *T, p funcs.Predicate[T]) *T {
	if tPtr == nil {
		return nil
	}

	if !p(*tPtr) {
		return nil
	}

	return tPtr
}
func FilterNot[T any](t *T, p funcs.Predicate[T]) *T { return Filter(t, predicate.Not(p)) }
