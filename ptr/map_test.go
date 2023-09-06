package ptr_test

import (
	"strconv"
	"testing"

	"github.com/mikhalytch/eggs/funcs/embed/unembed"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/ptr"
)

func TestMap(t *testing.T) {
	require.Nil(t, ptr.Map(ptr.None[string](), unembed.String[string]))
	require.Nil(t, ptr.Map(ptr.None[int](), strconv.Itoa))

	require.Equal(t, ptr.Of("abc"), ptr.Map(ptr.Of("abc"), unembed.String[string]))
	require.Equal(t, ptr.Of("123"), ptr.Map(ptr.Of(123), strconv.Itoa))
}
