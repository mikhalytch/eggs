package slices

import (
	"github.com/mikhalytch/eggs/funcs"
	"github.com/mikhalytch/eggs/opt"
)

func Map[V any, V1 any](mapper funcs.Mapper[V, V1]) funcs.Applier[[]V, []V1] {
	return func(vs []V) []V1 {
		res := make([]V1, 0, len(vs))

		for _, v := range vs {
			res = append(res, mapper(v))
		}

		return res
	}
}

func ToMapWithValues[K comparable, V any](valueGenerator func(i int, k K) V) funcs.Applier[[]K, map[K]V] {
	return func(ks []K) map[K]V {
		res := make(map[K]V, len(ks))

		for i, k := range ks {
			res[k] = valueGenerator(i, k)
		}

		return res
	}
}

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

func Head[T any](ts []T) T   { return ts[0] }
func Tail[T any](ts []T) []T { return ts[1:] }
func HeadOpt[T any](ts []T) opt.Option[T] {
	if len(ts) < 1 {
		return opt.None[T]()
	}

	return opt.Some(Head(ts))
}
