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

	require.True(t, try.Lazy[int](func() (int, error) { return 0, io.EOF }).IsFailure())
	require.False(t, try.Lazy[int](func() (int, error) { return 0, nil }).IsFailure())
}

func TestIsSuccess(t *testing.T) {
	require.False(t, try.Failure[string](io.EOF).IsSuccess())
	require.True(t, try.Success("abc").IsSuccess())

	require.False(t, try.Lazy[string](func() (string, error) { return "", io.EOF }).IsSuccess())
	require.True(t, try.Lazy[string](func() (string, error) { return "", nil }).IsSuccess())
}

func TestFailure_GetOrElse(t *testing.T) {
	require.Equal(t, 2, try.LiftOption(opt.None[int](), io.EOF).GetOrElse(2))
	require.Equal(t, "abc", try.Failure[string](io.EOF).GetOrElse("abc"))
}

func TestLazy_GetOrElse(t *testing.T) {
	require.Equal(t, 1, try.Lazy[int](func() (int, error) { return 1, nil }).GetOrElse(2))
	require.Equal(t, "abc", try.Lazy[string](func() (string, error) { return "", io.EOF }).GetOrElse("abc"))
}

func TestFailure_Get(t *testing.T) {
	r, err := try.Failure[string](io.EOF).Get()
	require.Equal(t, "", r)
	require.ErrorIs(t, err, io.EOF)
}

func TestLazy_Get(t *testing.T) {
	r, err := try.Lazy[string](func() (string, error) { return "", io.EOF }).Get()
	require.Equal(t, "", r)
	require.ErrorIs(t, err, io.EOF)
}

func TestSuccess_GetOrElse(t *testing.T) {
	require.Equal(t, 42, try.Success(42).GetOrElse(10))
}

func TestSuccess_Get(t *testing.T) {
	r, err := try.Success("abc").Get()
	require.NoError(t, err)
	require.Equal(t, "abc", r)
}

func TestForEach(t *testing.T) {
	tests := []struct {
		success, failure try.Try[any]
	}{
		{try.Success[any]("10"), try.Failure[any](io.EOF)},
		{try.Lazy[any](func() (any, error) { return "10", nil }), try.Lazy[any](func() (any, error) { return "", io.EOF })},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			cnt := 0
			inc := func(i any) { cnt++ }

			test.failure.ForEach(inc)
			require.Equal(t, 0, cnt)

			test.success.ForEach(inc)
			require.Equal(t, 1, cnt)
		})
	}
}

func TestToErr(t *testing.T) {
	tests := []struct {
		success try.Try[int]
		failure try.Try[any]
	}{
		{try.Success(10), try.Failure[any](io.EOF)},
		{try.Lazy[int](func() (int, error) { return 10, nil }), try.Lazy[any](func() (any, error) { return nil, io.EOF })},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			require.NoError(t, test.success.ToErr())
			require.Equal(t, io.EOF, test.failure.ToErr())
		})
	}
}

func TestToOpt(t *testing.T) {
	tests := []struct {
		success, failure try.Try[int]
	}{
		{try.Success(10), try.Failure[int](io.EOF)},
		{try.Lazy[int](func() (int, error) { return 10, nil }), try.Lazy[int](func() (int, error) { return 0, io.EOF })},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			require.NotEqual(t, opt.None[any](), test.failure.ToOpt())
			require.EqualValues(t, opt.None[any](), test.failure.ToOpt())
			require.Equal(t, opt.None[int](), test.failure.ToOpt())
			require.Equal(t, opt.Some(10), test.success.ToOpt())
		})
	}
}

func TestProc(t *testing.T) {
	f := func(i int) error { return io.EOF }
	s := func(i int) error { return nil }

	require.Equal(t, try.Failure[int](http.ErrMissingFile), try.Failure[int](http.ErrMissingFile).Proc(f))
	require.Equal(t, try.Failure[int](http.ErrMissingFile), try.Failure[int](http.ErrMissingFile).Proc(s))
	require.Equal(t, try.Failure[int](io.EOF), try.Success(1).Proc(f))
	require.Equal(t, try.Success(1), try.Success(1).Proc(s))
}

func TestProcFailure(t *testing.T) {
	cnt := 0
	f := func(e error) { cnt++ }

	try.Success(1).ProcFailure(f)
	require.Equal(t, 0, cnt)
	try.Failure[int](io.EOF).ProcFailure(f)
	require.Equal(t, 1, cnt)
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
	tests := []struct {
		newFailure func() try.Try[int]
	}{
		{func() try.Try[int] { return try.Failure[int](io.EOF) }},
		{func() try.Try[int] { return try.Lazy(func() (int, error) { return 0, io.EOF }) }},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			_, err := test.newFailure().FlatMap(try.Success[int]).Get()
			require.ErrorIs(t, err, io.EOF)

			_, err = test.newFailure().FlatMap(createFailureInt).Get()
			require.ErrorIs(t, err, io.EOF)
		})
	}
}

func TestSuccess_FlatMap(t *testing.T) {
	createFailureInt := func(_ int) try.Try[int] { return try.Failure[int](http.ErrAbortHandler) }

	tests := []struct {
		newSuccess func() try.Try[int]
	}{
		{func() try.Try[int] { return try.Success(1) }},
		{func() try.Try[int] { return try.Lazy[int](func() (int, error) { return 1, nil }) }},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f1, err := test.newSuccess().FlatMap(try.Success[int]).Get()
			require.NoError(t, err)
			require.Equal(t, 1, f1)

			_, err = test.newSuccess().FlatMap(createFailureInt).Get()
			require.ErrorIs(t, err, http.ErrAbortHandler)
		})
	}
}

func TestFlatMap(t *testing.T) {
	createSuccessInt := func(s string) try.Try[int] { return try.Trie(strconv.Atoi(s)) }

	createFailureString := func(_ int) try.Try[string] { return try.Failure[string](http.ErrAbortHandler) }

	t.Run("simple calc", func(t *testing.T) {
		scss, err := try.FlatMap(try.Success("1"), createSuccessInt).Get()
		require.NoError(t, err)
		require.Equal(t, 1, scss)

		_, err = try.FlatMap(try.Failure[string](io.EOF), createSuccessInt).Get()
		require.ErrorIs(t, err, io.EOF)

		_, err = try.FlatMap(try.Success(1), createFailureString).Get()
		require.ErrorIs(t, err, http.ErrAbortHandler)

		_, err = try.FlatMap(try.Failure[string](io.EOF), createSuccessInt).Get()
		require.ErrorIs(t, err, io.EOF)
	})
	t.Run("lazy calc", func(t *testing.T) {
		fm := try.FlatMap(try.Failure[string](io.EOF), func(r string) try.Try[int] { panic("") })
		_, err := fm.Get()
		require.ErrorIs(t, err, io.EOF)
	})
}

func TestLazy_Laziness(t *testing.T) {
	lazy := try.Lazy(func() (any, error) {
		panic("do")
	})

	lazy.Map(func(a any) any {
		return a
	})

	require.Panics(t, func() {
		lazy.IsFailure()
	})
}

func TestLazy_StackDepth(t *testing.T) {
	const testedDepth = 1_000_000

	createLazy := func() try.Try[int] {
		lazy := try.Lazy(func() (int, error) {
			return 0, nil
		})

		for i := 0; i < testedDepth; i++ {
			lazy = lazy.Map(func(i int) int {
				return i + 1
			})
		}

		lazy = lazy.Map(func(i int) int { return i * 2 })

		return lazy
	}

	t.Run("uncalled", func(t *testing.T) {
		createLazy()
	})
	t.Run("called", func(t *testing.T) {
		res, err := createLazy().Get()
		require.NoError(t, err)
		require.Equal(t, testedDepth*2, res)
	})
}
