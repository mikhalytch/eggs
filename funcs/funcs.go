package funcs

type (
	Applier[T any, R any] func(T) R

	Predicate[T any]        Applier[T, bool]
	Mapper[A any, B any]    Applier[A, B]
	OptMapper[A any, B any] Applier[A, *B]
)

func (f Applier[T, R]) Apply(t T) R { return f() }
