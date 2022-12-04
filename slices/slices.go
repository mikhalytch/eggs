package slices

import (
	"github.com/mikhalytch/eggs/funcs"
)

func Map[V any, V1 any](vs []V, mapper funcs.Mapper[V, V1]) []V1 {
	res := make([]V1, 0, len(vs))

	for _, v := range vs {
		res = append(res, mapper(v))
	}

	return res
}

func ToMapWithValues[K comparable, V any](ks []K, valueGenerator func(i int, k K) V) map[K]V {
	res := make(map[K]V, len(ks))

	for i, k := range ks {
		res[k] = valueGenerator(i, k)
	}

	return res
}

func Exists[T any](ts []T, predicate funcs.Predicate[T]) bool {
	for _, t := range ts {
		if predicate(t) {
			return true
		}
	}

	return false
}
