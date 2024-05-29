package deref

func Or[T any](ptr *T, or T) T {
	if ptr == nil {
		return or
	}

	return *ptr
}
