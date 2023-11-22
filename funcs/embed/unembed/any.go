package unembed

import baseConstraints "github.com/mikhalytch/eggs/constraints/base"

func String[T baseConstraints.String](t T) string { return string(t) }

func Bool[T baseConstraints.Bool](t T) bool { return bool(t) }

func Complex64[T baseConstraints.Complex64](t T) complex64    { return complex64(t) }
func Complex128[T baseConstraints.Complex128](t T) complex128 { return complex128(t) }

func Float64[T baseConstraints.Float64](t T) float64 { return float64(t) }
func Float32[T baseConstraints.Float32](t T) float32 { return float32(t) }

func Int[T baseConstraints.Int](t T) int       { return int(t) }
func Int64[T baseConstraints.Int64](t T) int64 { return int64(t) }
func Int32[T baseConstraints.Int32](t T) int32 { return int32(t) }
func Int16[T baseConstraints.Int16](t T) int16 { return int16(t) }
func Int8[T baseConstraints.Int8](t T) int8    { return int8(t) }

func Uint[T baseConstraints.Uint](t T) uint          { return uint(t) }
func Uint8[T baseConstraints.Uint8](t T) uint8       { return uint8(t) }
func Uint16[T baseConstraints.Uint16](t T) uint16    { return uint16(t) }
func Uint32[T baseConstraints.Uint32](t T) uint32    { return uint32(t) }
func Uint64[T baseConstraints.Uint64](t T) uint64    { return uint64(t) }
func Uintptr[T baseConstraints.Uintptr](t T) uintptr { return uintptr(t) }
