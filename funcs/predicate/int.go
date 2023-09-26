package predicate

import (
	"github.com/mikhalytch/eggs/constraints"
)

func NonZero[A constraints.Integer](a A) bool { return a != 0 }
