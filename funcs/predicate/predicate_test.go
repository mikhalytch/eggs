package predicate_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/predicate"
)

func TestNonZero(t *testing.T) {
	require.True(t, predicate.NonZero(1))
	require.True(t, predicate.NonZero(int8(-2)))
	require.True(t, predicate.NonZero(uint(20)))
	require.False(t, predicate.NonZero(uint(0)))
	require.False(t, predicate.NonZero(uint16(0)))
}

func TestAlways(t *testing.T) {
	require.True(t, predicate.Always(1))
	require.True(t, predicate.Always("abc"))
	require.True(t, predicate.Always[any](nil))
}

func TestNot(t *testing.T) {
	require.False(t, predicate.Not(predicate.Always[int])(1))
	require.False(t, predicate.Not(predicate.Always[string])("abc"))
	require.False(t, predicate.Not(predicate.Always[any])(nil))
}

func TestNever(t *testing.T) {
	require.False(t, predicate.Never(1))
	require.False(t, predicate.Never("abc"))
	require.False(t, predicate.Never[any](nil))
}
