package maps_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/predicate"
	"github.com/mikhalytch/eggs/maps"
)

func TestFilterValues(t *testing.T) {
	var n map[string]any

	require.Len(t, maps.FilterValues(n, predicate.Always[any]), 0)

	require.Len(t, maps.FilterValues(map[string]string{}, predicate.Always[string]), 0)
	require.Equal(t, map[string]int{"b": 1}, maps.FilterValues(map[string]int{"a": 0, "b": 1}, predicate.NonZero[int]))
}
