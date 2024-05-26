package reembed_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/embed/reembed"
)

func TestReembedInt(t *testing.T) {
	t.Run("int32", func(t *testing.T) {
		type (
			iTestType uint16
			oTestType int32
		)

		testVal := iTestType(123)

		emb := oTestType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, reembed.Int32[iTestType, oTestType](testVal))
	})
}
