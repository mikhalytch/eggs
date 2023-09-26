package strconv

import (
	"strconv"

	"github.com/mikhalytch/eggs/constraints"
)

func UtoA[T constraints.Unsigned](t T) string { return strconv.FormatUint(uint64(t), 10) }
