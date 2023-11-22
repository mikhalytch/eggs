package strconv

import (
	"strconv"

	baseConstraints "github.com/mikhalytch/eggs/constraints/base"
)

func UtoA[T baseConstraints.Unsigned](t T) string { return strconv.FormatUint(uint64(t), 10) }
