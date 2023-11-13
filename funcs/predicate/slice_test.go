package predicate_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/predicate"
)

func TestContains(t *testing.T) {
	require.False(t, predicate.Contains([]string{"a", "b"})("c"))
	require.True(t, predicate.Contains([]int{1, 2})(2))
	require.True(t, predicate.Not(predicate.Contains([]int{1, 2}))(3))
}
