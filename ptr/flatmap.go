package ptr

import (
	"github.com/mikhalytch/eggs/funcs"
)

func FlatMap[T any, R any](t *T, fMapper funcs.Mapper[*T, *R]) *R {
	if t == nil {
		return nil
	}

	res := fMapper(t)

	return res
}
