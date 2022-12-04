package predicate

import (
	"golang.org/x/exp/constraints"
)

func NonZero[A constraints.Integer](a A) bool { return a != 0 }
func Always[T any](T) bool                    { return true }
func Never[T any](T) bool                     { return false }
