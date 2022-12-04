package strconv

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func UtoA[T constraints.Unsigned](t T) string { return strconv.FormatUint(uint64(t), 10) }
