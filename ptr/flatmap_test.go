package ptr_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/embed/unembed"
	"github.com/mikhalytch/eggs/ptr"
)

func TestFlatMap(t *testing.T) {
	fmSS := func(t *string) *string { return ptr.Of(unembed.String[string](*t)) }
	fmIS := func(t *int) *string { return ptr.Of(strconv.Itoa(*t)) }

	require.Nil(t, ptr.FlatMap(ptr.None[string](), fmSS))
	require.Nil(t, ptr.FlatMap(ptr.None[int](), fmIS))

	require.Equal(t, ptr.Of("abc"), ptr.FlatMap(ptr.Of("abc"), fmSS))
	require.Equal(t, ptr.Of("123"), ptr.FlatMap(ptr.Of(123), fmIS))
}
