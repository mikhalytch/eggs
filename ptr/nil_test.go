package ptr_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/ptr"
)

func TestNone(t *testing.T) {
	require.Equal(t, (*string)(nil), ptr.None[string]())
	require.Equal(t, (*int)(nil), ptr.None[int]())
}
