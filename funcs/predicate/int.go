package predicate

import (
	"golang.org/x/exp/constraints"
)

func NonZero[A constraints.Integer](a A) bool { return a != 0 }
