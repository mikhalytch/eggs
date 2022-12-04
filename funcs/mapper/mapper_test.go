package mapper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/mapper"
)

func TestString(t *testing.T) {
	t.Run("Name", func(t *testing.T) {
		type Name string
		name := Name("abc")
		require.NotEqual(t, "abc", name)
		require.Equal(t, "abc", mapper.String(name))
	})
}

func TestIdentity(t *testing.T) {
	require.Equal(t, "a", mapper.Identity("a"))
	require.Equal(t, 1, mapper.Identity(1))
}
