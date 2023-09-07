package unembed_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/embed/unembed"
)

func TestUnEmbedString(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		type testType string
		testVal := "abc"

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.String[testType](emb))
	})
}

func TestUnEmbedBool(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		type testType bool
		testVal := false

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Bool[testType](emb))
	})
}

func TestUnEmbedInt(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		type testType int
		testVal := 123

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Int[testType](emb))
	})
	t.Run("int64", func(t *testing.T) {
		type testType int64
		testVal := int64(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Int64[testType](emb))
	})
	t.Run("int32", func(t *testing.T) {
		type testType int32
		testVal := int32(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Int32[testType](emb))
	})
	t.Run("int16", func(t *testing.T) {
		type testType int16
		testVal := int16(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Int16[testType](emb))
	})
	t.Run("int8", func(t *testing.T) {
		type testType int8
		testVal := int8(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Int8[testType](emb))
	})
}

func TestUnEmbedUint(t *testing.T) {
	t.Run("Uint8", func(t *testing.T) {
		type testType uint8
		testVal := uint8(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Uint8[testType](emb))
	})
	t.Run("Uint16", func(t *testing.T) {
		type testType uint16
		testVal := uint16(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Uint16[testType](emb))
	})
	t.Run("Uint32", func(t *testing.T) {
		type testType uint32
		testVal := uint32(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Uint32[testType](emb))
	})
	t.Run("Uint64", func(t *testing.T) {
		type testType uint64
		testVal := uint64(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Uint64[testType](emb))
	})
	t.Run("Uintptr", func(t *testing.T) {
		type testType uintptr
		testVal := uintptr(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Uintptr[testType](emb))
	})
	t.Run("Uint", func(t *testing.T) {
		type testType uint
		testVal := uint(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Uint[testType](emb))
	})
}

func TestUnEmbedFloat(t *testing.T) {
	t.Run("Float32", func(t *testing.T) {
		type testType float32
		testVal := float32(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Float32[testType](emb))
	})
	t.Run("Float64", func(t *testing.T) {
		type testType float64
		testVal := float64(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Float64[testType](emb))
	})
}

func TestUnEmbedComplex(t *testing.T) {
	t.Run("Complex64", func(t *testing.T) {
		type testType complex64
		testVal := complex(float32(123), 345)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Complex64[testType](emb))
	})
	t.Run("Complex128", func(t *testing.T) {
		type testType complex128
		testVal := complex(123, 456)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, testVal, unembed.Complex128[testType](emb))
	})
}
