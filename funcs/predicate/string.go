package predicate

import (
	"strings"

	baseConstraints "github.com/mikhalytch/eggs/constraints/base"
	"github.com/mikhalytch/eggs/funcs/embed/unembed"
)

func IsBlank[T baseConstraints.String](t T) bool {
	return len(strings.TrimSpace(unembed.String(t))) == 0
}
