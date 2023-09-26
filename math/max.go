package math

import (
	"github.com/mikhalytch/eggs/constraints"
)

func Max[T constraints.Ordered](first T, other ...T) T {
	maxF := func(a, b T) T {
		if a > b {
			return a
		}

		return b
	}

	res := first
	for _, _b := range other {
		res = maxF(res, _b)
	}

	return res
}
