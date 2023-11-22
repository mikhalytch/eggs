package constraints

type (
	String interface{ ~string }

	Bool interface{ ~bool }

	Complex64  interface{ ~complex64 }
	Complex128 interface{ ~complex128 }
	Complex    interface {
		Complex64 | Complex128
	}

	Float64 interface{ ~float64 }
	Float32 interface{ ~float32 }
	Float   interface {
		Float32 | Float64
	}

	Int    interface{ ~int }
	Int64  interface{ ~int64 }
	Int32  interface{ ~int32 }
	Int16  interface{ ~int16 }
	Int8   interface{ ~int8 }
	Signed interface {
		Int | Int64 | Int32 | Int16 | Int8
	}

	Uint     interface{ ~uint }
	Uint8    interface{ ~uint8 }
	Uint16   interface{ ~uint16 }
	Uint32   interface{ ~uint32 }
	Uint64   interface{ ~uint64 }
	Uintptr  interface{ ~uintptr }
	Unsigned interface {
		Uint | Uint8 | Uint16 | Uint32 | Uint64 | Uintptr
	}

	Integer interface {
		Signed | Unsigned
	}
	Ordered interface {
		Integer | Float | String
	}
)
