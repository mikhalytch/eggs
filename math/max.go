package math

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](first T, other ...T) T {
	max := func(a, b T) T {
		if a > b {
			return a
		}

		return b
	}

	res := first
	for _, _b := range other {
		res = max(res, _b)
	}

	return res
}
