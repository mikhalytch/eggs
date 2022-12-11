package math

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](first T, other ...T) T {
	min := func(a, b T) T {
		if a < b {
			return a
		}

		return b
	}

	res := first
	for _, _b := range other {
		res = min(res, _b)
	}

	return res
}
