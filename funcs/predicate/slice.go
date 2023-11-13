package predicate

import (
	"golang.org/x/exp/slices"

	"github.com/mikhalytch/eggs/funcs"
)

func Contains[E ~[]T, T comparable](s E) funcs.Predicate[T] {
	return func(t T) bool {
		return slices.Contains(s, t)
	}
}
