package unembed

import "github.com/mikhalytch/eggs/constraints"

func String[T constraints.String](t T) string { return string(t) }

func Complex64[T constraints.Complex64](t T) complex64    { return complex64(t) }
func Complex128[T constraints.Complex128](t T) complex128 { return complex128(t) }

func Float64[T constraints.Float64](t T) float64 { return float64(t) }
func Float32[T constraints.Float32](t T) float32 { return float32(t) }

func Int[T constraints.Int](t T) int       { return int(t) }
func Int64[T constraints.Int64](t T) int64 { return int64(t) }
func Int32[T constraints.Int32](t T) int32 { return int32(t) }
func Int16[T constraints.Int16](t T) int16 { return int16(t) }
func Int8[T constraints.Int8](t T) int8    { return int8(t) }

func Uint[T constraints.Uint](t T) uint          { return uint(t) }
func Uint8[T constraints.Uint8](t T) uint8       { return uint8(t) }
func Uint16[T constraints.Uint16](t T) uint16    { return uint16(t) }
func Uint32[T constraints.Uint32](t T) uint32    { return uint32(t) }
func Uint64[T constraints.Uint64](t T) uint64    { return uint64(t) }
func Uintptr[T constraints.Uintptr](t T) uintptr { return uintptr(t) }
