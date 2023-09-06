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
