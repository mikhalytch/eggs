package maps

import (
	"github.com/mikhalytch/eggs/tuple"
)

func Entries[M ~map[K]V, K comparable, V any](m M) []tuple.Tuple[K, V] {
	res := make([]tuple.Tuple[K, V], 0, len(m))

	for k, v := range m {
		res = append(res, tuple.Of(k, v))
	}

	return res
}
