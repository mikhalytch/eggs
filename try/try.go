package try

import (
	"github.com/mikhalytch/eggs/funcs"
	"github.com/mikhalytch/eggs/opt"
)

type (
	// Try is like Either, with left type simplified to error
	// (Basically, Try[T].ToEither() :: Either[error, T]).
	Try[R any] interface {
		IsSuccess() bool

		// OrElse returns contained value if this is success; otherwise returns given `other` value.
		OrElse(other R) R

		// ForEach executes given procedure for success, skips for failure.
		ForEach(procedure funcs.Procedure[R]) Try[R]

		// ToErr is nil for success, and error for failure
		ToErr() error

		// ToOpt converts success to opt.Some, failure to opt.None.
		ToOpt() opt.Option[R]

		// ----- Proc/Map/FlatMap -----

		// Proc does nothing for left; it turns success to failure on proc error
		Proc(proc funcs.FallibleFunction[R]) Try[R]
		ProcFailure(proc funcs.Procedure[error]) Try[R]

		// Map is category unchanging method, variant of Map function.
		Map(mapper funcs.Mapper[R, R]) Try[R]
		MapFailure(lm funcs.Mapper[error, error]) Try[R]

		// FlatMap is category unchanging method, variant of FlatMap function.
		FlatMap(fMap FMapper[R, R]) Try[R]
	}
	// success represents the success value of Try.
	success[R any] struct{ r R }
	// failure represents failure value of Try.
	failure[R any] struct{ err error }

	FMapper[R, V any] funcs.Mapper[R, Try[V]]
)

// ----- success -----

func (r success[R]) IsSuccess() bool                                { return true }
func (r success[R]) OrElse(_ R) R                                   { return r.r }
func (r success[R]) ToErr() error                                   { return nil }
func (r success[R]) ToOpt() opt.Option[R]                           { return opt.Some(r.r) }
func (r success[R]) Proc(proc funcs.FallibleFunction[R]) Try[R]     { return Trie(r.r, proc(r.r)) }
func (r success[R]) ProcFailure(_ funcs.Procedure[error]) Try[R]    { return r }
func (r success[R]) Map(m funcs.Mapper[R, R]) Try[R]                { return Map[R](r, m) }
func (r success[R]) MapFailure(_ funcs.Mapper[error, error]) Try[R] { return r }
func (r success[R]) FlatMap(fMap FMapper[R, R]) Try[R]              { return FlatMap[R](r, fMap) }
func (r success[R]) ForEach(p funcs.Procedure[R]) Try[R] {
	p(r.r)

	return r
}

// ----- failure -----

func (l failure[R]) IsSuccess() bool                         { return false }
func (l failure[R]) OrElse(other R) R                        { return other }
func (l failure[R]) ToErr() error                            { return l.err }
func (l failure[R]) ToOpt() opt.Option[R]                    { return opt.None[R]() }
func (l failure[R]) Proc(_ funcs.FallibleFunction[R]) Try[R] { return l }
func (l failure[R]) ProcFailure(proc funcs.Procedure[error]) Try[R] {
	proc(l.err)

	return l
}
func (l failure[R]) Map(m funcs.Mapper[R, R]) Try[R]                 { return Map[R](l, m) }
func (l failure[R]) MapFailure(lm funcs.Mapper[error, error]) Try[R] { return Failure[R](lm(l.err)) }
func (l failure[R]) FlatMap(fMap FMapper[R, R]) Try[R]               { return FlatMap[R](l, fMap) }
func (l failure[R]) ForEach(_ funcs.Procedure[R]) Try[R]             { return l }

// ----- General -----

func Success[R any](r R) Try[R]       { return success[R]{r: r} }
func Failure[R any](err error) Try[R] { return failure[R]{err: err} }
func Trie[R any](r R, err error) Try[R] {
	if err != nil {
		return Failure[R](err)
	}

	return Success(r)
}

func LiftOption[R any](o opt.Option[R], err error) Try[R] {
	/*Trie(o.Get()).MapFailure(mapper.Always(err))*/
	return opt.Fold(o, Success[R], func() Try[R] { return Failure[R](err) })
}

// Map returns success containing the result of applying f to m if it's not failure;
// otherwise it returns unchanged failure.
func Map[R, V any](e Try[R], f funcs.Mapper[R, V]) Try[V] {
	if s, ok := e.(success[R]); ok {
		return success[V]{r: f(s.r)}
	}

	return Failure[V](e.ToErr())
}

// FlatMap returns success with result of applying f to R, if this is success; otherwise it returns unchanged failure.
func FlatMap[R, V any](e Try[R], f FMapper[R, V]) Try[V] {
	if s, ok := e.(success[R]); ok {
		return f(s.r)
	}

	return Failure[V](e.ToErr())
}
