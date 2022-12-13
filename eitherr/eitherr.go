package eitherr

import (
	"github.com/mikhalytch/eggs/funcs"
	"github.com/mikhalytch/eggs/opt"
)

type (
	/*
		Either represents value that can be of one of two types, left or right.
		With that said, left should be generic over some L, right - over R:
			type Either[L, R any] interface {}
			type right[R any] struct { R }
			type left[L any] struct { L }
		However, staying on a practical side, and taking into account following facts:
		  1. left type L can be simplified to error
		  2. golang type system makes it no simple using complex generic types
		  3. things may change in the future, and left may want to reclaim its type
		Either is simplified to
			type Eitherr[R any] interface {}
			type right[R any] struct { R }
			type left[R any] struct { error }
		with all auxiliary functions type signatures became simplified as well.
		(However, there will be no Swap function)
	*/
	Eitherr[R any] interface {
		// OrElse returns contained value if this is right; otherwise returns given `other` value.
		OrElse(other R) R

		// ForEach executes given procedure for right, skips for left.
		ForEach(procedure funcs.Procedure[R]) Eitherr[R]

		// ToErr is nil for right, and error for left
		ToErr() error

		// ToOpt converts right to opt.Some, left to opt.None.
		ToOpt() opt.Option[R]

		// Map is category unchanging method, variant of Map function.
		Map(mapper funcs.Mapper[R, R]) Eitherr[R]
		MapLeft(lm funcs.Mapper[error, error]) Eitherr[R]

		// FlatMap is category unchanging method, variant of FlatMap function.
		FlatMap(fMap FMapper[R, R]) Eitherr[R]
	}
	// right represents the right value of Eitherr.
	right[R any] struct{ r R }
	// left represents left value of Eitherr.
	left[R any] struct{ err error }

	FMapper[R, V any] funcs.Mapper[R, Eitherr[V]]
)

// ----- right -----

func (r right[R]) OrElse(_ R) R                                    { return r.r }
func (r right[R]) ToErr() error                                    { return nil }
func (r right[R]) ToOpt() opt.Option[R]                            { return opt.Some(r.r) }
func (r right[R]) Map(m funcs.Mapper[R, R]) Eitherr[R]             { return Map[R, R](r, m) }
func (r right[R]) MapLeft(_ funcs.Mapper[error, error]) Eitherr[R] { return r }
func (r right[R]) FlatMap(fMap FMapper[R, R]) Eitherr[R]           { return FlatMap[R, R](r, fMap) }
func (r right[R]) ForEach(p funcs.Procedure[R]) Eitherr[R] {
	p(r.r)

	return r
}

// ----- left -----

func (l left[R]) OrElse(other R) R                                 { return other }
func (l left[R]) ToErr() error                                     { return l.err }
func (l left[R]) ToOpt() opt.Option[R]                             { return opt.None[R]() }
func (l left[R]) Map(m funcs.Mapper[R, R]) Eitherr[R]              { return Map[R, R](l, m) }
func (l left[R]) MapLeft(lm funcs.Mapper[error, error]) Eitherr[R] { return Left[R](lm(l.err)) }
func (l left[R]) FlatMap(fMap FMapper[R, R]) Eitherr[R]            { return FlatMap[R, R](l, fMap) }
func (l left[R]) ForEach(_ funcs.Procedure[R]) Eitherr[R]          { return l }

// ----- General -----

func Right[R any](r R) Eitherr[R]      { return right[R]{r: r} }
func Left[R any](err error) Eitherr[R] { return left[R]{err: err} }
func FromFallible[R any](r R, err error) Eitherr[R] {
	if err != nil {
		return Left[R](err)
	}

	return Right[R](r)
}

func LiftOption[R any](o opt.Option[R], err error) Eitherr[R] {
	/*FromFallible(o.Get()).MapLeft(mapper.Always(err))*/
	return opt.Fold[R, Eitherr[R]](o, Right[R], func() Eitherr[R] { return Left[R](err) })
}

// Map returns right containing the result of applying f to m if it's not left; otherwise it returns unchanged left.
func Map[R, V any](e Eitherr[R], f funcs.Mapper[R, V]) Eitherr[V] {
	if s, ok := e.(right[R]); ok {
		return right[V]{r: f(s.r)}
	}

	return Left[V](e.ToErr())
}

// FlatMap returns right with result of applying f to R, if this is right; otherwise it returns unchanged left.
func FlatMap[R, V any](e Eitherr[R], f FMapper[R, V]) Eitherr[V] {
	if s, ok := e.(right[R]); ok {
		return f(s.r)
	}

	return Left[V](e.ToErr())
}
