package mapper

import (
	"github.com/mikhalytch/eggs/funcs"
)

func Identity[A any](a A) A                { return a }
func ToAny[A any](a A) any                 { return a }
func Always[A any](a A) funcs.Mapper[A, A] { return func(_ A) A { return a } }
func Struct[A any](_ A) struct{}           { return struct{}{} }
