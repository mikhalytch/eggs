package predicate

import (
	baseConstraints "github.com/mikhalytch/eggs/constraints/base"
)

func NonZero[A baseConstraints.Integer](a A) bool { return a != 0 }
