package dflt_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/dflt"
)

func TestOf(t *testing.T) {
	require.Equal(t, 0, dflt.Of[int]())
	require.Equal(t, "", dflt.Of[string]())
	require.Equal(t, (*int)(nil), dflt.Of[*int]())
}
