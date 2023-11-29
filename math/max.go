package math

import (
	baseConstraints "github.com/mikhalytch/eggs/constraints/base"
	"github.com/mikhalytch/eggs/ptr"
	"github.com/mikhalytch/eggs/slices"
)

func Max[T baseConstraints.Ordered](first T, other ...T) T {
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

func MaxOpt[T baseConstraints.Ordered](ts ...*T) *T {
	elems := slices.Map(ptr.ToSlice[T])(ts)

	flat := slices.Join(elems...)
	hOpt := slices.HeadOpt(flat)

	return ptr.Map(hOpt, func(h T) T { return Max(h, slices.Tail(flat)...) })
}
