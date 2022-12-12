package mapkeys

import (
	"github.com/mikhalytch/eggs/funcs"
)

// Map returns m keys mapped with mapper.
func Map[M ~map[K]V, K comparable, V any, K1 comparable](m M, mapper funcs.Mapper[K, K1]) []K1 {
	res := make([]K1, 0, len(m))

	for k := range m {
		res = append(res, mapper(k))
	}

	return res
}
