package strconv

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func StoA[T constraints.Signed](t T) string { return strconv.FormatInt(int64(t), 10) }
