package maps_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/maps"
	"github.com/mikhalytch/eggs/tuple"
)

func TestEntries(t *testing.T) {
	var n map[string]any

	require.Len(t, maps.Entries(n), 0)

	require.Len(t, maps.Entries(map[string]int{}), 0)

	require.EqualValues(t, []tuple.Tuple[int, string]{tuple.Of(1, "a")}, maps.Entries(map[int]string{1: "a"}))
}
