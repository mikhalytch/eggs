package slices_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/mapper"
	"github.com/mikhalytch/eggs/slices"
	"github.com/mikhalytch/eggs/strconv"
)

func TestMap(t *testing.T) {
	require.Equal(t, []int{}, slices.Map(mapper.Identity[int])([]int{}))
	require.Equal(t, []int{1, 2, 3}, slices.Map(mapper.Identity[int])([]int{1, 2, 3}))
	require.Equal(t, []string{}, slices.Map(strconv.StoA[int])([]int{}))
	require.Equal(t, []string{"1", "2", "2"}, slices.Map(strconv.StoA[int])([]int{1, 2, 2}))
	require.Equal(t, []any{'a', 'b'}, slices.Map(mapper.ToAny[rune])([]rune{'a', 'b'}))
}

func TestFlatMap(t *testing.T) {
	singleFMap := slices.FlatMap[int](func(i int) []int { return []int{i} })
	require.Equal(t, []int{}, singleFMap.Apply([]int{}))
	require.Equal(t, []int{1, 2}, singleFMap.Apply([]int{1, 2}))
	require.Equal(t, []int{}, singleFMap.Apply(nil))

	emptyFMap := slices.FlatMap[int](func(_ int) []int { return []int{} })
	require.Equal(t, []int{}, emptyFMap.Apply([]int{1, 2, 3}))
	require.Equal(t, []int{}, emptyFMap.Apply([]int{}))
	require.Equal(t, []int{}, emptyFMap.Apply(nil))

	nilFMap := slices.FlatMap[int](func(_ int) []int { return nil })
	require.Equal(t, []int{}, nilFMap.Apply([]int{1, 2, 3}))
	require.Equal(t, []int{}, nilFMap.Apply([]int{}))
	require.Equal(t, []int{}, nilFMap.Apply(nil))

	doubleFMap := slices.FlatMap[string](func(i string) []string { return []string{i, i} })
	require.Equal(t, []string{}, doubleFMap.Apply([]string{}))
	require.Equal(t, []string{"a", "a", "bb", "bb"}, doubleFMap.Apply([]string{"a", "bb"}))
	require.Equal(t, []string{}, doubleFMap.Apply(nil))
}

func TestToMapWithValues(t *testing.T) {
	require.Equal(t, map[string]int{}, slices.ToMapWithValues(func(k string) int { return 0 })([]string{}))
	require.Equal(t, map[string]int{"a": 0, "b": 0},
		slices.ToMapWithValues(func(k string) int { return 0 })([]string{"a", "b"}))
	require.Equal(t, map[string]int{"a": 0, "b": 1},
		slices.ToMapWithValues(func(k string) int { return int(k[0] - 'a') })([]string{"a", "b"}))
}

func TestToMapWithKeys(t *testing.T) {
	require.Equal(t, map[string]int{}, slices.ToMapWithKeys(func(v int) string { return "" })([]int{}))
	require.Equal(t, map[string]int{"a": 1},
		slices.ToMapWithKeys(func(v int) string { return "a" })([]int{0, 1}))
	require.Equal(t, map[string]int{"a": 0, "b": 1},
		slices.ToMapWithKeys(func(v int) string { return string([]byte{byte(v + 'a')}) })([]int{0, 1}))
}
