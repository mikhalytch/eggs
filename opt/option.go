package opt

import (
	"errors"

	"github.com/mikhalytch/eggs/deref"
	"github.com/mikhalytch/eggs/funcs"
)

type (
	/*
		Option represents the simplest monad that can contain some value, available for mapping with Map,
			or nothing (none).
		With that said, Map mapper will never face nil.
	*/
	Option[T any] interface {
		IsDefined() bool
		IsEmpty() bool
		Get() (T, error)
		OrElse(other T) T
		ToSlice() []T
		Filter(p funcs.Predicate[T]) Option[T]
		Foreach(p funcs.Procedure[T]) Option[T]

		// Map method can't change the category; see Map function for that matter.
		Map(mapper funcs.Mapper[T, T]) Option[T]
		// FlatMap method can't change the category; see FlatMap function for that matter.
		FlatMap(mapper FMapper[T, T]) Option[T]
	}
	some[T any] struct{ t T }
	none[T any] struct{}

	FMapper[T, V any] funcs.Mapper[T, Option[V]]
)

func OfPtr[T any](t *T) Option[*T] {
	if t == nil {
		return none[*T]{}
	}

	return some[*T]{t: t}
}

func None[T any]() Option[T]    { return none[T]{} }
func Some[T any](t T) Option[T] { return some[T]{t: t} }

// ----- none -----

var ErrNoneGet = errors.New("none.Get")

func (n none[T]) IsDefined() bool                        { return false }
func (n none[T]) IsEmpty() bool                          { return !n.IsDefined() }
func (n none[T]) Get() (T, error)                        { return deref.Of(new(T)), ErrNoneGet }
func (n none[T]) OrElse(other T) T                       { return other }
func (n none[T]) ToSlice() []T                           { return []T{} }
func (n none[T]) Filter(_ funcs.Predicate[T]) Option[T]  { return n }
func (n none[T]) Map(m funcs.Mapper[T, T]) Option[T]     { return Map[T, T](n, m) }
func (n none[T]) FlatMap(m FMapper[T, T]) Option[T]      { return FlatMap[T, T](n, m) }
func (n none[T]) Foreach(_ funcs.Procedure[T]) Option[T] { return n }

// ----- some -----

func (s some[T]) IsDefined() bool { return true }
func (s some[T]) IsEmpty() bool   { return !s.IsDefined() }
func (s some[T]) Get() (T, error) { return s.t, nil }
func (s some[T]) OrElse(_ T) T    { return s.t }
func (s some[T]) ToSlice() []T    { return []T{s.t} }
func (s some[T]) Filter(p funcs.Predicate[T]) Option[T] {
	if p(s.t) {
		return s
	}

	return None[T]()
}
func (s some[T]) Map(m funcs.Mapper[T, T]) Option[T] { return Map[T, T](s, m) }
func (s some[T]) FlatMap(m FMapper[T, T]) Option[T]  { return FlatMap[T, T](s, m) }
func (s some[T]) Foreach(p funcs.Procedure[T]) Option[T] {
	p(s.t)

	return s
}

// ----- Map/FlatMap -----

func Map[T, V any](m Option[T], f funcs.Mapper[T, V]) Option[V] {
	if s, ok := m.(some[T]); ok {
		return Some[V](f(s.t))
	}

	return None[V]()
}

func FlatMap[T, V any](m Option[T], f FMapper[T, V]) Option[V] {
	if s, ok := m.(some[T]); ok {
		return f(s.t)
	}

	return None[V]()
}
