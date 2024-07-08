package slices

import (
	"github.com/mikhalytch/eggs/funcs"
)

type (
	FMapper[K, K1 any] funcs.Mapper[K, []K1]
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

func FlatMap[V, V1 any](fMap FMapper[V, V1]) funcs.Applier[[]V, []V1] {
	return func(vs []V) []V1 {
		res := make([]V1, 0)

		for _, v := range vs {
			res = append(res, fMap(v)...)
		}

		return res
	}
}

func ToMapWithValues[K comparable, V any](valueGenerator func(k K) V) funcs.Applier[[]K, map[K]V] {
	return func(ks []K) map[K]V {
		res := make(map[K]V, len(ks))

		for _, k := range ks {
			res[k] = valueGenerator(k)
		}

		return res
	}
}

func ToMapWithKeys[K comparable, V any](keyProducer func(V) K) funcs.Applier[[]V, map[K]V] {
	return func(vs []V) map[K]V {
		res := make(map[K]V, 0)

		for _, v := range vs {
			res[keyProducer(v)] = v
		}

		return res
	}
}

func ToMultimapWithKeys[K comparable, V any](keyProducer func(V) K) funcs.Applier[[]V, map[K][]V] {
	return func(vs []V) map[K][]V {
		res := make(map[K][]V, 0)

		for _, v := range vs {
			key := keyProducer(v)

			ex := res[key]
			ex = append(ex, v)
			res[key] = ex
		}

		return res
	}
}
