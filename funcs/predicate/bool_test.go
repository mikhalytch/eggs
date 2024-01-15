package predicate_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/predicate"
)

func TestIsTrue(t *testing.T) {
	type bb bool

	tr, fa := bb(true), bb(false)

	require.True(t, predicate.IsTrue(tr))
	require.False(t, predicate.IsTrue(fa))
}

func TestIsFalse(t *testing.T) {
	type bb bool

	tr, fa := bb(true), bb(false)

	require.True(t, predicate.IsFalse(fa))
	require.False(t, predicate.IsFalse(tr))
}
