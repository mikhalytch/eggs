package dflt

func Of[T any]() T { return *new(T) }
