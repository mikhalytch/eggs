package tuple_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/tuple"
)

func TestTuple_Swap(t *testing.T) {
	require.Equal(t, tuple.Of(1, "abc"), tuple.Of("abc", 1).Swap())
}

func TestTuple_T1(t *testing.T) {
	require.Equal(t, 1, tuple.Err(1, nil).T1())
	require.Equal(t, nil, tuple.FromFallible("abc", nil).Swap().T1())
}

func TestTuple_T2(t *testing.T) {
	type Name string

	require.Equal(t, Name("some"), tuple.Of(io.EOF, Name("some")).T2())
}
