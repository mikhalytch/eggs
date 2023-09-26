package math

import (
	"github.com/mikhalytch/eggs/constraints"
)

func Min[T constraints.Ordered](first T, other ...T) T {
	minF := func(a, b T) T {
		if a < b {
			return a
		}

		return b
	}

	res := first
	for _, _b := range other {
		res = minF(res, _b)
	}

	return res
}
