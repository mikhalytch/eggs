package predicate

import (
	"github.com/mikhalytch/eggs/funcs"
)

func Always[T any](T) bool { return true }
func Never[T any](T) bool  { return false }
func Not[T any](p funcs.Predicate[T]) funcs.Predicate[T] {
	return func(t T) bool {
		return !p(t)
	}
}
