package tuple

type (
	Tuple[A, B any] struct {
		a A
		b B
	}
)

// -----

func (t Tuple[A, B]) Swap() Tuple[B, A] { return Tuple[B, A]{a: t.b, b: t.a} }
func (t Tuple[A, B]) T1() A             { return t.a }
func (t Tuple[A, B]) T2() B             { return t.b }

// -----

func Of[A, B any](a A, b B) Tuple[A, B]                  { return Tuple[A, B]{a: a, b: b} }
func FromFallible[A any](a A, err error) Tuple[A, error] { return Of(a, err) }
func Err[A any](a A, err error) Tuple[A, error]          { return FromFallible(a, err) }
