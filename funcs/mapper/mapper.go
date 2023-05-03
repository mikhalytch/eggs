package mapper

import (
	"github.com/mikhalytch/eggs/funcs"
)

func String[A ~string](a A) string         { return string(a) }
func Identity[A any](a A) A                { return a }
func ToAny[A any](a A) any                 { return a }
func Always[A any](a A) funcs.Mapper[A, A] { return func(_ A) A { return a } }
