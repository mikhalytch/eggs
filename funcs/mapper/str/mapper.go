package str

import (
	baseConstraints "github.com/mikhalytch/eggs/constraints/base"
	"github.com/mikhalytch/eggs/funcs"
)

func Prepend[A baseConstraints.String](p A) funcs.Mapper[A, A] { return func(a A) A { return p + a } }
