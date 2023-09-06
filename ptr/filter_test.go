package ptr_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/ptr"
)

func TestFilterNot(t *testing.T) {
	evenInt := func(i int) bool { return i%2 == 0 }

	require.Nil(t, nil, ptr.FilterNot(nil, evenInt))
	require.Equal(t, ptr.Nil[int](), ptr.FilterNot(ptr.Of(2), evenInt))
	require.Equal(t, ptr.Of(1), ptr.FilterNot(ptr.Of(1), evenInt))
}
