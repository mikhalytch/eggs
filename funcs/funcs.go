package funcs

type (
	Function0[V any]         func() V
	FallibleFunction0[V any] func() (V, error)
	Procedure[T any]         func(T)
	Routine                  func()
	FallibleRoutine          func() error

	Function[T, V any]      func(T) V
	Predicate[T any]        Function[T, bool]
	FallibleFunction[T any] Function[T, error]
	Mapper[A any, B any]    Function[A, B]

	Applier[T any, R any] Function[T, R]
)

func (f Applier[T, R]) Apply(t T) R { return f(t) }
