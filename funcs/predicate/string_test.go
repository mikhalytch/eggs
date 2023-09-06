package predicate_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/predicate"
)

func TestIsBlank(t *testing.T) {
	require.True(t, predicate.IsBlank(""))
	require.True(t, predicate.IsBlank("  "))
	require.True(t, predicate.IsBlank("	"))
	require.True(t, predicate.IsBlank(`

`))

	type testType string

	testVal := testType("")

	require.True(t, predicate.IsBlank(testVal))
	require.False(t, predicate.IsBlank(testVal+"sdf"))
	require.False(t, predicate.IsBlank(" sdf"))
}
