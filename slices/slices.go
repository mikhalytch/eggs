package slices

import (
	"github.com/mikhalytch/eggs/funcs"
	"github.com/mikhalytch/eggs/opt"
)

func Exists[T any](predicate funcs.Predicate[T]) funcs.Applier[[]T, bool] {
	return func(ts []T) bool {
		for _, t := range ts {
			if predicate(t) {
				return true
			}
		}

		return false
	}
}

func Filter[T any](predicate funcs.Predicate[T]) funcs.Applier[[]T, []T] {
	return func(ts []T) []T {
		res := make([]T, 0)

		for _, t := range ts {
			if predicate(t) {
				res = append(res, t)
			}
		}

		return res
	}
}

// Split returns funcs.Applier producing suitable for predicate values on the left, and non-suitable on the right.
func Split[T any](predicate funcs.Predicate[T]) func([]T) ([]T, []T) {
	return func(ts []T) ([]T, []T) {
		suitable, nonSuitable := make([]T, 0), make([]T, 0)

		for _, t := range ts {
			if predicate(t) {
				suitable = append(suitable, t)
			} else {
				nonSuitable = append(nonSuitable, t)
			}
		}

		return suitable, nonSuitable
	}
}

func Head[T any](ts []T) T   { return ts[0] }
func Tail[T any](ts []T) []T { return ts[1:] }
func HeadOption[T any](ts []T) opt.Option[T] {
	if len(ts) < 1 {
		return opt.None[T]()
	}

	return opt.Some(Head(ts))
}

func HeadOpt[T any](ts []T) *T {
	if len(ts) < 1 {
		return nil
	}

	return &ts[0]
}

func Flatten[T any](tss [][]T) []T { return Join(tss...) }
func Join[T any](tss ...[]T) []T {
	if len(tss) == 0 {
		return nil
	}

	res := make([]T, 0)
	for _, ts := range tss {
		res = append(res, ts...)
	}

	return res
}
