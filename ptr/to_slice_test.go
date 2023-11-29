package ptr_test

import (
	"testing"

	"github.com/mikhalytch/eggs/ptr"

	"github.com/stretchr/testify/require"
)

func TestToSlice(t *testing.T) {
	require.Empty(t, ptr.ToSlice[int](nil))
	require.Equal(t, []string{"abc"}, ptr.ToSlice[string](ptr.Of("abc")))
}
