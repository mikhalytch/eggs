package try_test

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/mapper"
	"github.com/mikhalytch/eggs/opt"
	"github.com/mikhalytch/eggs/try"
)

func TestIsFailure(t *testing.T) {
	require.True(t, try.Failure[int](io.EOF).IsFailure())
	require.False(t, try.Success(1).IsFailure())
}

func TestIsSuccess(t *testing.T) {
	require.False(t, try.Failure[string](io.EOF).IsSuccess())
	require.True(t, try.Success("abc").IsSuccess())
}

func TestFailure_OrElse(t *testing.T) {
	require.Equal(t, 2, try.LiftOption(opt.None[int](), io.EOF).OrElse(2))
	require.Equal(t, "abc", try.Failure[string](io.EOF).OrElse("abc"))
}

func TestSuccess_OrElse(t *testing.T) {
	require.Equal(t, 42, try.Success[int](42).OrElse(10))
}

func TestForEach(t *testing.T) {
	cnt := 0
	inc := func(i any) { cnt++ }

	try.Failure[any](io.EOF).ForEach(inc)
	require.Equal(t, 0, cnt)

	try.Success[any]("10").ForEach(inc)
	require.Equal(t, 1, cnt)
}

func TestToErr(t *testing.T) {
	require.NoError(t, try.Success(10).ToErr())
	require.Equal(t, io.EOF, try.Failure[any](io.EOF).ToErr())
}

func TestToOpt(t *testing.T) {
	require.NotEqual(t, opt.None[int](), try.Failure[any](io.EOF).ToOpt())
	require.EqualValues(t, opt.None[int](), try.Failure[any](io.EOF).ToOpt())
	require.Equal(t, opt.None[int](), try.Failure[int](io.EOF).ToOpt())
	require.Equal(t, opt.Some(10), try.Success(10).ToOpt())
}

func TestProc(t *testing.T) {
	f := func(i int) error { return io.EOF }
	s := func(i int) error { return nil }

	require.Equal(t, try.Failure[int](http.ErrMissingFile), try.Failure[int](http.ErrMissingFile).Proc(f))
	require.Equal(t, try.Failure[int](http.ErrMissingFile), try.Failure[int](http.ErrMissingFile).Proc(s))
	require.Equal(t, try.Failure[int](io.EOF), try.Success(1).Proc(f))
	require.Equal(t, try.Success(1), try.Success(1).Proc(s))
}

func TestFailure_Map(t *testing.T) {
	require.Equal(t, try.Failure[int](io.EOF), try.Failure[int](io.EOF).Map(mapper.Identity[int]))
	require.Equal(t, try.Failure[int](io.EOF), try.Failure[int](io.EOF).Map(func(i int) int { return i * 1000 }))
}

func TestSuccess_Map(t *testing.T) {
	require.Equal(t, try.Success(10), try.Success(10).Map(mapper.Identity[int]))
	require.Equal(t, try.Success(10), try.Success(1).Map(func(i int) int { return i * 10 }))
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
	require.Equal(t, try.Success(1), try.Map(try.Trie(fallible()), mapper.Identity[int]))
	require.Equal(t, try.Failure[int](io.EOF), try.Map(try.Trie(fallible()), mapper.Identity[int]))
	require.Equal(t, try.Success(30),
		try.Map(try.Trie(fallible()), func(i int) int { return i * 10 }))
	require.Equal(t, try.Failure[int](io.EOF),
		try.Map(try.Trie(fallible()), func(i int) int { return i * 10 }))
}

func TestMapFailure(t *testing.T) {
	alw := mapper.Always(http.ErrMissingFile)

	require.Equal(t, try.Success(1), try.Success(1).MapFailure(alw))
	require.Equal(t, try.Failure[int](http.ErrMissingFile), try.Failure[int](io.EOF).MapFailure(alw))

	cond := func(e1 error) error {
		if errors.Is(e1, io.EOF) {
			return http.ErrAbortHandler
		}

		return http.ErrBodyNotAllowed
	}
	require.Equal(t, try.Failure[int](http.ErrAbortHandler), try.Failure[int](io.EOF).MapFailure(cond))
	require.Equal(t, try.Failure[int](http.ErrBodyNotAllowed), try.Failure[int](http.ErrMissingFile).MapFailure(cond))
}

func TestFailure_FlatMap(t *testing.T) {
	createFailureInt := func(_ int) try.Try[int] { return try.Failure[int](http.ErrAbortHandler) }

	require.Equal(t, try.Failure[int](io.EOF), try.Failure[int](io.EOF).FlatMap(try.Success[int]))
	require.Equal(t, try.Failure[int](io.EOF), try.Failure[int](io.EOF).FlatMap(createFailureInt))
}

func TestSuccess_FlatMap(t *testing.T) {
	createFailureInt := func(_ int) try.Try[int] { return try.Failure[int](http.ErrAbortHandler) }

	require.Equal(t, try.Success(1), try.Success[int](1).FlatMap(try.Success[int]))
	require.Equal(t, try.Failure[int](http.ErrAbortHandler), try.Success[int](1).FlatMap(createFailureInt))
}

func TestFlatMap(t *testing.T) {
	createSuccessInt := func(s string) try.Try[int] { return try.Trie(strconv.Atoi(s)) }
	createFailureString := func(_ int) try.Try[string] { return try.Failure[string](http.ErrAbortHandler) }

	require.Equal(t, try.Success(1), try.FlatMap(try.Success[string]("1"), createSuccessInt))
	require.Equal(t, try.Failure[int](io.EOF), try.FlatMap(try.Failure[string](io.EOF), createSuccessInt))

	require.Equal(t, try.Failure[string](http.ErrAbortHandler),
		try.FlatMap(try.Success[int](1), createFailureString))
	require.Equal(t, try.Failure[int](io.EOF), try.FlatMap(try.Failure[string](io.EOF), createSuccessInt))
}
