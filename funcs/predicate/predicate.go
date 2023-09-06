package predicate

import (
	"golang.org/x/exp/constraints"

	"github.com/mikhalytch/eggs/funcs"
)

func NonZero[A constraints.Integer](a A) bool { return a != 0 }
func Always[T any](T) bool                    { return true }
func Never[T any](T) bool                     { return false }

func Not[T any](p funcs.Predicate[T]) funcs.Predicate[T] {
	return func(t T) bool {
		return !p(t)
	}
}
