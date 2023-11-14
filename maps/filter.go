package maps

import (
	"github.com/mikhalytch/eggs/funcs"
)

func FilterValues[M ~map[K]V, K comparable, V any](m M, vp funcs.Predicate[V]) M {
	res := make(M)

	for k, v := range m {
		if vp(v) {
			res[k] = v
		}
	}

	return res
}
