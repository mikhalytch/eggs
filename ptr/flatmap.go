package ptr

type (
	FlatMapper[T any, R any] func(T) *R
)

func FlatMap[T any, R any](t *T, fMapper FlatMapper[T, R]) *R {
	if t == nil {
		return nil
	}

	res := fMapper(*t)

	return res
}
