package predicate

import (
	"strings"

	"github.com/mikhalytch/eggs/constraints"
	"github.com/mikhalytch/eggs/funcs/embed/unembed"
)

func IsBlank[T constraints.String](t T) bool {
	return len(strings.TrimSpace(unembed.String(t))) == 0
}
