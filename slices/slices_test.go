package slices_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/mapper"
	"github.com/mikhalytch/eggs/funcs/predicate"
	"github.com/mikhalytch/eggs/opt"
	"github.com/mikhalytch/eggs/slices"
	"github.com/mikhalytch/eggs/strconv"
)

func TestMap(t *testing.T) {
	require.Equal(t, []int{}, slices.Map(mapper.Identity[int])([]int{}))
	require.Equal(t, []int{1, 2, 3}, slices.Map(mapper.Identity[int])([]int{1, 2, 3}))
	require.Equal(t, []string{}, slices.Map(strconv.StoA[int])([]int{}))
	require.Equal(t, []string{"1", "2", "2"}, slices.Map(strconv.StoA[int])([]int{1, 2, 2}))
}

func TestFlatMap(t *testing.T) {
	singleFMap := slices.FlatMap(func(i int) []int { return []int{i} })
	require.Equal(t, []int{}, singleFMap.Apply([]int{}))
	require.Equal(t, []int{1, 2}, singleFMap.Apply([]int{1, 2}))
	require.Equal(t, []int{}, singleFMap.Apply(nil))

	emptyFMap := slices.FlatMap(func(_ int) []int { return []int{} })
	require.Equal(t, []int{}, emptyFMap.Apply([]int{1, 2, 3}))
	require.Equal(t, []int{}, emptyFMap.Apply([]int{}))
	require.Equal(t, []int{}, emptyFMap.Apply(nil))

	nilFMap := slices.FlatMap(func(_ int) []int { return nil })
	require.Equal(t, []int{}, nilFMap.Apply([]int{1, 2, 3}))
	require.Equal(t, []int{}, nilFMap.Apply([]int{}))
	require.Equal(t, []int{}, nilFMap.Apply(nil))

	doubleFMap := slices.FlatMap(func(i string) []string { return []string{i, i} })
	require.Equal(t, []string{}, doubleFMap.Apply([]string{}))
	require.Equal(t, []string{"a", "a", "bb", "bb"}, doubleFMap.Apply([]string{"a", "bb"}))
	require.Equal(t, []string{}, doubleFMap.Apply(nil))
}

func TestToMapWithValues(t *testing.T) {
	require.Equal(t, map[string]int{}, slices.ToMapWithValues(func(i int, k string) int { return 0 })([]string{}))
	require.Equal(t, map[string]int{"a": 0, "b": 0},
		slices.ToMapWithValues(func(i int, k string) int { return 0 })([]string{"a", "b"}))
	require.Equal(t, map[string]int{"a": 0, "b": 1},
		slices.ToMapWithValues(func(i int, k string) int { return i })([]string{"a", "b"}))
}

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
