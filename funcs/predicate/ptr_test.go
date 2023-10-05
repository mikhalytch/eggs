package predicate_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/predicate"
	"github.com/mikhalytch/eggs/ptr"
	"github.com/mikhalytch/eggs/slices"
)

func TestNonNil(t *testing.T) {
	l := []*int{ptr.Of(1), nil, ptr.Of(2)}
	got := slices.Filter(predicate.NonNil[int])(l)
	require.Len(t, got, 2)
	require.Equal(t, 1, *got[0])
	require.Equal(t, 2, *got[1])
}

func TestNil(t *testing.T) {
	l := []*int{ptr.Of(1), nil, ptr.Of(2)}
	got := slices.Filter(predicate.Nil[int])(l)
	require.Len(t, got, 1)
	require.Nil(t, got[0])
}
