package strconv

import (
	"strconv"

	"github.com/mikhalytch/eggs/constraints"
)

func StoA[T constraints.Signed](t T) string { return strconv.FormatInt(int64(t), 10) }
