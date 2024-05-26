package embed_test

import (
	strconv2 "strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mikhalytch/eggs/funcs/embed"
)

func TestEmbedString(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		type testType string

		testVal := "abc"

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.String[testType](testVal))
	})
}

func TestEmbedBool(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		type testType bool

		tests := []struct {
			v bool
		}{
			{true},
			{false},
		}

		for i, test := range tests {
			t.Run(strconv2.Itoa(i), func(t *testing.T) {
				testVal := test.v

				emb := testType(testVal)
				require.NotEqual(t, testVal, emb)

				require.Equal(t, emb, embed.Bool[testType](testVal))
			})
		}
	})
}

func TestEmbedInt(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		type testType int

		testVal := 123

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Int[testType](testVal))
	})
	t.Run("int64", func(t *testing.T) {
		type testType int64

		testVal := int64(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Int64[testType](testVal))
	})
	t.Run("int32", func(t *testing.T) {
		type testType int32

		testVal := int32(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Int32[testType](testVal))
	})
	t.Run("int16", func(t *testing.T) {
		type testType int16

		testVal := int16(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Int16[testType](testVal))
	})
	t.Run("int8", func(t *testing.T) {
		type testType int8

		testVal := int8(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Int8[testType](testVal))
	})
}

func TestEmbedUint(t *testing.T) {
	t.Run("Uint8", func(t *testing.T) {
		type testType uint8

		testVal := uint8(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Uint8[testType](testVal))
	})
	t.Run("Uint16", func(t *testing.T) {
		type testType uint16

		testVal := uint16(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Uint16[testType](testVal))
	})
	t.Run("Uint32", func(t *testing.T) {
		type testType uint32

		testVal := uint32(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Uint32[testType](testVal))
	})
	t.Run("Uint64", func(t *testing.T) {
		type testType uint64

		testVal := uint64(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Uint64[testType](testVal))
	})
	t.Run("Uintptr", func(t *testing.T) {
		type testType uintptr

		testVal := uintptr(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Uintptr[testType](testVal))
	})
	t.Run("Uint", func(t *testing.T) {
		type testType uint

		testVal := uint(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Uint[testType](testVal))
	})
}

func TestEmbedFloat(t *testing.T) {
	t.Run("Float32", func(t *testing.T) {
		type testType float32

		testVal := float32(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Float32[testType](testVal))
	})
	t.Run("Float64", func(t *testing.T) {
		type testType float64

		testVal := float64(123)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Float64[testType](testVal))
	})
}

func TestEmbedComplex(t *testing.T) {
	t.Run("Complex64", func(t *testing.T) {
		type testType complex64

		testVal := complex(float32(123), 345)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Complex64[testType](testVal))
	})
	t.Run("Complex128", func(t *testing.T) {
		type testType complex128

		testVal := complex(123, 456)

		emb := testType(testVal)
		require.NotEqual(t, testVal, emb)

		require.Equal(t, emb, embed.Complex128[testType](testVal))
	})
}
