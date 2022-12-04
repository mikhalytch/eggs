package ptr_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/ptr"
)

func TestOf(t *testing.T) {
	require.Equal(t, 1, *ptr.Of(1))
	require.Equal(t, "a", *ptr.Of("a"))

	type Name string

	require.Equal(t, Name("abc"), *ptr.Of(Name("abc")))
	require.Equal(t, Name("abc"), *ptr.Of[Name]("abc"))
}
