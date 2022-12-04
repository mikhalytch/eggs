package funcs

type (
	Predicate[T any]        func(T) bool
	Mapper[A any, B any]    func(A) B
	OptMapper[A any, B any] func(A) *B
)
