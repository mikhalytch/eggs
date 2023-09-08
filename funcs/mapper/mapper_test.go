package mapper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/mapper"
	"github.com/mikhalytch/eggs/slices"
)

func TestIdentity(t *testing.T) {
	require.Equal(t, "a", mapper.Identity("a"))
	require.Equal(t, 1, mapper.Identity(1))
}

func TestToAny(t *testing.T) {
	require.Equal(t, "a", mapper.ToAny("a"))
	require.Equal(t, 1, mapper.ToAny(1))
}

func TestAlways(t *testing.T) {
	alw := mapper.Always(5)

	require.Equal(t, 5, alw(-100))
	require.Equal(t, 5, alw(0))
	require.Equal(t, 5, alw(1))
}

func TestStruct(t *testing.T) {
	l := []string{"a", "b"}

	index := slices.ToMapWithValues(mapper.Struct[string])(l)

	_, aOk := index["a"]
	_, bOk := index["b"]
	_, cOk := index["c"]

	require.True(t, aOk)
	require.True(t, bOk)
	require.False(t, cOk)
}
