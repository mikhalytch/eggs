package mapvalues

import (
	"github.com/mikhalytch/eggs/funcs"
)

func Filter[M ~map[K]V, K comparable, V any](m M, predicate funcs.Predicate[V]) []V {
	res := make([]V, 0, len(m))

	for _, v := range m {
		if predicate(v) {
			res = append(res, v)
		}
	}

	return res
}
