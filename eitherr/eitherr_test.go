package eitherr_test

import (
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/eitherr"
	"github.com/mikhalytch/eggs/funcs/mapper"
	"github.com/mikhalytch/eggs/opt"
)

func TestLeft_OrElse(t *testing.T) {
	require.Equal(t, 2, eitherr.LiftOption(opt.None[int](), io.EOF).OrElse(2))
	require.Equal(t, "abc", eitherr.Left[string](io.EOF).OrElse("abc"))
}

func TestRight_OrElse(t *testing.T) {
	require.Equal(t, 42, eitherr.Right[int](42).OrElse(10))
}

func TestForEach(t *testing.T) {
	cnt := 0
	inc := func(i any) { cnt++ }

	eitherr.Left[any](io.EOF).ForEach(inc)
	require.Equal(t, 0, cnt)

	eitherr.Right[any]("10").ForEach(inc)
	require.Equal(t, 1, cnt)
}

func TestToErr(t *testing.T) {
	require.NoError(t, eitherr.Right(10).ToErr())
	require.Equal(t, io.EOF, eitherr.Left[any](io.EOF).ToErr())
}

func TestToOpt(t *testing.T) {
	require.NotEqual(t, opt.None[int](), eitherr.Left[any](io.EOF).ToOpt())
	require.EqualValues(t, opt.None[int](), eitherr.Left[any](io.EOF).ToOpt())
	require.Equal(t, opt.None[int](), eitherr.Left[int](io.EOF).ToOpt())
	require.Equal(t, opt.Some(10), eitherr.Right(10).ToOpt())
}

func TestLeft_Map(t *testing.T) {
	require.Equal(t, eitherr.Left[int](io.EOF), eitherr.Left[int](io.EOF).Map(mapper.Identity[int]))
	require.Equal(t, eitherr.Left[int](io.EOF), eitherr.Left[int](io.EOF).Map(func(i int) int { return i * 1000 }))
}

func TestRight_Map(t *testing.T) {
	require.Equal(t, eitherr.Right(10), eitherr.Right(10).Map(mapper.Identity[int]))
	require.Equal(t, eitherr.Right(10), eitherr.Right(1).Map(func(i int) int { return i * 10 }))
}

func TestMap(t *testing.T) {
	cnt := 0
	fallible := func() (int, error) {
		cnt++
		if cnt%2 == 0 {
			return 0, io.EOF
		}

		return cnt, nil
	}
	require.Equal(t, eitherr.Right(1), eitherr.Map(eitherr.FromFallible(fallible()), mapper.Identity[int]))
	require.Equal(t, eitherr.Left[int](io.EOF), eitherr.Map(eitherr.FromFallible(fallible()), mapper.Identity[int]))
	require.Equal(t, eitherr.Right(30),
		eitherr.Map(eitherr.FromFallible(fallible()), func(i int) int { return i * 10 }))
	require.Equal(t, eitherr.Left[int](io.EOF),
		eitherr.Map(eitherr.FromFallible(fallible()), func(i int) int { return i * 10 }))
}

func TestLeft_FlatMap(t *testing.T) {
	createLeftInt := func(_ int) eitherr.Eitherr[int] { return eitherr.Left[int](http.ErrAbortHandler) }

	require.Equal(t, eitherr.Left[int](io.EOF), eitherr.Left[int](io.EOF).FlatMap(eitherr.Right[int]))
	require.Equal(t, eitherr.Left[int](io.EOF), eitherr.Left[int](io.EOF).FlatMap(createLeftInt))
}

func TestRight_FlatMap(t *testing.T) {
	createLeftInt := func(_ int) eitherr.Eitherr[int] { return eitherr.Left[int](http.ErrAbortHandler) }

	require.Equal(t, eitherr.Right(1), eitherr.Right[int](1).FlatMap(eitherr.Right[int]))
	require.Equal(t, eitherr.Left[int](http.ErrAbortHandler), eitherr.Right[int](1).FlatMap(createLeftInt))
}

func TestFlatMap(t *testing.T) {
	createRightInt := func(s string) eitherr.Eitherr[int] { return eitherr.FromFallible(strconv.Atoi(s)) }
	createLeftString := func(_ int) eitherr.Eitherr[string] { return eitherr.Left[string](http.ErrAbortHandler) }

	require.Equal(t, eitherr.Right(1), eitherr.FlatMap(eitherr.Right[string]("1"), createRightInt))
	require.Equal(t, eitherr.Left[int](io.EOF), eitherr.FlatMap(eitherr.Left[string](io.EOF), createRightInt))

	require.Equal(t, eitherr.Left[string](http.ErrAbortHandler),
		eitherr.FlatMap(eitherr.Right[int](1), createLeftString))
	require.Equal(t, eitherr.Left[int](io.EOF), eitherr.FlatMap(eitherr.Left[string](io.EOF), createRightInt))
}
