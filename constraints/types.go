package constraints

type (
	String interface{ ~string }

	Complex64  interface{ ~complex64 }
	Complex128 interface{ ~complex128 }

	Float64 interface{ ~float64 }
	Float32 interface{ ~float32 }

	Int   interface{ ~int }
	Int64 interface{ ~int64 }
	Int32 interface{ ~int32 }
	Int16 interface{ ~int16 }
	Int8  interface{ ~int8 }

	Uint    interface{ ~uint }
	Uint8   interface{ ~uint8 }
	Uint16  interface{ ~uint16 }
	Uint32  interface{ ~uint32 }
	Uint64  interface{ ~uint64 }
	Uintptr interface{ ~uintptr }
)

// AnyX as named at
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md

type (
	AnyString = String

	AnyComplex64  = Complex64
	AnyComplex128 = Complex128

	AnyFloat64 = Float64
	AnyFloat32 = Float32

	AnyInt   = Int
	AnyInt64 = Int64
	AnyInt32 = Int32
	AnyInt16 = Int16
	AnyInt8  = Int8

	AnyUint    = Uint
	AnyUint8   = Uint8
	AnyUint16  = Uint16
	AnyUint32  = Uint32
	AnyUint64  = Uint64
	AnyUintptr = Uintptr
)
