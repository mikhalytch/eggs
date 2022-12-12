package mapkeys

import (
	"github.com/mikhalytch/eggs/funcs"
	"github.com/mikhalytch/eggs/opt"
)

// Map returns m keys mapped with mapper.
func Map[M ~map[K]V, K comparable, V any, K1 comparable](m M, mapper funcs.Mapper[K, K1]) []K1 {
	res := make([]K1, 0, len(m))

	for k := range m {
		res = append(res, mapper(k))
	}

	return res
}

// FlatMap returns m keys mapped with fMap; in case fMap(key) is none, key is ignored.
func FlatMap[M ~map[K]V, K comparable, V any, K1 comparable](m M, fMap opt.FMapper[K, K1]) []K1 {
	res := make([]K1, 0, len(m))

	for k := range m {
		fMap(k).Foreach(func(k K1) {
			res = append(res, k)
		})
	}

	return res
}
