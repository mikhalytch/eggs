package maps

import (
	"github.com/mikhalytch/eggs/funcs"
)

func MapKeys[M ~map[K]V, K comparable, V any, K1 comparable](m M, mapper funcs.Mapper[K, K1]) map[K1]V {
	res := make(map[K1]V, len(m))
	for k, v := range m {
		res[mapper(k)] = v
	}

	return res
}

func MapValues[M ~map[K]V, K comparable, V any, V1 any](m M, mapper funcs.Mapper[V, V1]) map[K]V1 {
	res := make(map[K]V1, len(m))
	for k, v := range m {
		res[k] = mapper(v)
	}

	return res
}
