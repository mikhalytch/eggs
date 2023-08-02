package slices_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/predicate"
	"github.com/mikhalytch/eggs/opt"
	"github.com/mikhalytch/eggs/slices"
)

func TestExists(t *testing.T) {
	t.Run("empty slice always results in false", func(t *testing.T) {
		require.False(t, slices.Exists(predicate.Never[int])([]int{}))
		require.False(t, slices.Exists(predicate.Always[int])([]int{}))
	})
	t.Run("different predicates", func(t *testing.T) {
		require.True(t, slices.Exists(predicate.Always[int])([]int{1}))
		require.False(t, slices.Exists(predicate.Never[int])([]int{1}))

		require.True(t, slices.Exists(func(a int) bool { return a == 1 })([]int{1}))
		require.False(t, slices.Exists(func(a int) bool { return a == 1 })([]int{2}))
	})
}

func TestFilter(t *testing.T) {
	require.Equal(t, []int{1, 2}, slices.Filter(func(i int) bool { return i < 3 })([]int{1, 2, 3, 4, 5}))
	require.Equal(t, []string{"abc", "cde"},
		slices.Filter(func(i string) bool { return i < "efg" })([]string{"abc", "cde", "efg"}))
}

func TestSplit(t *testing.T) {
	even, odd := slices.Split(func(i int) bool { return i%2 == 0 })(nil)
	require.Equal(t, []int{}, even)
	require.Equal(t, []int{}, odd)

	even, odd = slices.Split(func(i int) bool { return i%2 == 0 })([]int{})
	require.Equal(t, []int{}, even)
	require.Equal(t, []int{}, odd)

	even, odd = slices.Split(func(i int) bool { return i%2 == 0 })([]int{0, 1, 2})
	require.Equal(t, []int{0, 2}, even)
	require.Equal(t, []int{1}, odd)

	as, bs := slices.Split(func(i rune) bool { return i == 'a' })([]rune{'a', 'b'})
	require.Equal(t, []rune{'a'}, as)
	require.Equal(t, []rune{'b'}, bs)
}

func TestHead(t *testing.T) {
	t.Run("existing elements", func(t *testing.T) {
		require.Equal(t, 1, slices.Head([]int{1, 2, 3}))
		require.Equal(t, "first", slices.Head([]string{"first", "second"}))
	})
	t.Run("panics on empty slice", func(t *testing.T) {
		require.Panics(t, func() { slices.Head([]int{}) })
		require.Panics(t, func() { slices.Head[int](nil) })
	})
}

func TestTail(t *testing.T) {
	t.Run("slice has enough elements", func(t *testing.T) {
		require.Equal(t, []int{2, 3}, slices.Tail([]int{1, 2, 3}))
		require.Equal(t, []string{}, slices.Tail([]string{"first"}))
		require.Equal(t, []string{"second", "third"}, slices.Tail([]string{"first", "second", "third"}))
	})
	t.Run("panics on empty slice", func(t *testing.T) {
		require.Panics(t, func() { slices.Tail([]int{}) })
		type Name string
		require.Panics(t, func() { slices.Tail[Name](nil) })
	})
}

func TestHeadOpt(t *testing.T) {
	require.Equal(t, opt.Some(1), slices.HeadOpt([]int{1, 2}))
	require.Equal(t, opt.Some("a"), slices.HeadOpt([]string{"a"}))
	require.Equal(t, opt.None[int](), slices.HeadOpt([]int{}))
}

func TestFlatten(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		require.Equal(t, ([]int)(nil), slices.Flatten[int](nil))
		require.Equal(t, ([]int)(nil), slices.Flatten[int]([][]int{}))
		require.Equal(t, []int{1, 2}, slices.Flatten([][]int{{1, 2}}))
		require.Equal(t, []int{1, 2, 1, 2, 3}, slices.Flatten([][]int{{1, 2}, {1, 2}, {3}}))
	})
	t.Run("strings", func(t *testing.T) {
		require.Equal(t, ([]string)(nil), slices.Flatten[string](nil))
		require.Equal(t, ([]string)(nil), slices.Flatten[string]([][]string{}))
		require.Equal(t, []string{"a", "b"}, slices.Flatten([][]string{{"a", "b"}}))
		require.Equal(t, []string{"a", "b", "a", "b", "abc"},
			slices.Flatten([][]string{{"a", "b"}, {"a", "b"}, {"abc"}}))
	})
}
