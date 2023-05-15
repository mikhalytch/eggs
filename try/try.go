package try

import (
	goLazy "github.com/hymkor/go-lazy"

	"github.com/mikhalytch/eggs/dflt"
	"github.com/mikhalytch/eggs/funcs"
	"github.com/mikhalytch/eggs/opt"
)

type (
	// Try is like Either, with left type simplified to error
	// (i.e. Try[T].ToEither() === Either[error, T]).
	Try[R any] interface {
		Is[R]
		To[R]
		Monadic[R]
		GoIdiomatic[R]
	}
	Is[R any] interface {
		IsFailure() bool
		IsSuccess() bool
	}
	To[R any] interface {
		// ToErr is nil for success, and error for failure
		ToErr() error

		// ToOpt converts success to opt.Some, failure to opt.None.
		ToOpt() opt.Option[R]
	}
	Monadic[R any] interface {
		// GetOrElse returns contained value if this is success; otherwise returns given `other` value.
		GetOrElse(other R) R
		// ForEach executes given procedure for success, skips for failure.
		ForEach(procedure funcs.Procedure[R]) Try[R]
		Functional[R]
	}
	Functional[R any] interface {
		// Proc does nothing for left; it turns success to failure on proc error
		Proc(proc funcs.FallibleFunction[R]) Try[R]
		ProcFailure(proc funcs.Procedure[error]) Try[R]

		// Map is category unchanging method, variant of Map function.
		Map(mapper funcs.Mapper[R, R]) Try[R]
		MapFailure(lm funcs.Mapper[error, error]) Try[R]

		// FlatMap is category unchanging method, variant of FlatMap function.
		FlatMap(fMap FMapper[R, R]) Try[R]
	}
	GoIdiomatic[R any] interface {
		Get() (R, error)
	}
	// success represents the success value of Try.
	success[R any] struct{ r R }
	// failure represents failure value of Try.
	failure[R any] struct{ err error }

	// lazy represents lazy computation.
	lazy[R any] struct{ delayed *goLazy.Of[Try[R]] }

	FMapper[R, V any] funcs.Mapper[R, Try[V]]
)

// ----- success -----

func (r success[R]) IsFailure() bool                                { return !r.IsSuccess() }
func (r success[R]) IsSuccess() bool                                { return true }
func (r success[R]) GetOrElse(_ R) R                                { return r.r }
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
func (r success[R]) Get() (R, error) { return r.r, nil }

// ----- failure -----

func (l failure[R]) IsFailure() bool                         { return true }
func (l failure[R]) IsSuccess() bool                         { return !l.IsFailure() }
func (l failure[R]) GetOrElse(other R) R                     { return other }
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
func (l failure[R]) Get() (R, error)                                 { return dflt.Of[R](), l.err }

// ----- lazy -----

func (l lazy[R]) IsFailure() bool      { return l.delayed.Value().IsFailure() }
func (l lazy[R]) IsSuccess() bool      { return l.delayed.Value().IsSuccess() }
func (l lazy[R]) ToErr() error         { return l.delayed.Value().ToErr() }
func (l lazy[R]) ToOpt() opt.Option[R] { return l.delayed.Value().ToOpt() }
func (l lazy[R]) GetOrElse(other R) R  { return l.delayed.Value().GetOrElse(other) }
func (l lazy[R]) ForEach(procedure funcs.Procedure[R]) Try[R] {
	return l.delayed.Value().ForEach(procedure)
}

func (l lazy[R]) Proc(proc funcs.FallibleFunction[R]) Try[R] {
	return Lazy(func() (R, error) {
		return l.delayed.Value().Proc(proc).Get()
	})
}

func (l lazy[R]) ProcFailure(proc funcs.Procedure[error]) Try[R] {
	return l.delayed.Value().ProcFailure(proc)
}

func (l lazy[R]) Map(mapper funcs.Mapper[R, R]) Try[R] {
	return Lazy(func() (R, error) {
		return l.delayed.Value().Map(mapper).Get()
	})
}

func (l lazy[R]) MapFailure(lm funcs.Mapper[error, error]) Try[R] {
	return l.delayed.Value().MapFailure(lm)
}

func (l lazy[R]) FlatMap(fMap FMapper[R, R]) Try[R] { return l.delayed.Value().FlatMap(fMap) }
func (l lazy[R]) Get() (R, error)                   { return l.delayed.Value().Get() }

// ----- General -----

func Success[R any](r R) Try[R]       { return success[R]{r: r} }
func Failure[R any](err error) Try[R] { return failure[R]{err: err} }

func From[R any](r R, err error) Try[R] { return Trie(r, err) } // alias to Trie
func Trie[R any](r R, err error) Try[R] {
	if err != nil {
		return Failure[R](err)
	}

	return Success(r)
}

func Lazy[R any](thunk funcs.FallibleFunction0[R]) Try[R] {
	return lazy[R]{delayed: goLazy.New(func() Try[R] { return From(thunk()) })}
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
func FlatMap[R, V any](e Try[R], fMapper FMapper[R, V]) Try[V] {
	return Lazy[V](func() (V, error) {
		r, err := e.Get()
		if err != nil {
			return dflt.Of[V](), err
		}

		return fMapper(r).Get()
	})
}
