package mapvalues_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/predicate"
	"github.com/mikhalytch/eggs/maps/mapvalues"
)

func TestFilter(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		require.ElementsMatch(t, []string{}, mapvalues.Filter(map[int]string{}, func(a string) bool { return a != "a" }))
		require.ElementsMatch(t, []string{"b"},
			mapvalues.Filter(map[int]string{1: "a", 2: "b"}, func(a string) bool { return a != "a" }))
	})
	t.Run("Int", func(t *testing.T) {
		require.ElementsMatch(t, []int{}, mapvalues.Filter(map[string]int{}, predicate.NonZero[int]))
		require.ElementsMatch(t, []int{1},
			mapvalues.Filter(map[string]int{"a": 2, "b": 1}, func(a int) bool { return a != 2 }))
	})
}
