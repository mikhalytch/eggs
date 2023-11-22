package strconv

import (
	"strconv"

	baseConstraints "github.com/mikhalytch/eggs/constraints/base"
)

func StoA[T baseConstraints.Signed](t T) string { return strconv.FormatInt(int64(t), 10) }
