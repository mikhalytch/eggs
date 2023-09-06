package ptr

import (
	"github.com/mikhalytch/eggs/funcs"
)

func Map[T any, R any](t *T, mapper funcs.Mapper[T, R]) *R {
	if t == nil {
		return nil
	}

	return Of(mapper(*t))
}
