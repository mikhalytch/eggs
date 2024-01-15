package predicate

import constraints "github.com/mikhalytch/eggs/constraints/base"

func IsTrue[T constraints.Bool](t T) bool  { return bool(t) }
func IsFalse[T constraints.Bool](t T) bool { return !bool(t) }
