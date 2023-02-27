package utils

func Map[T, U any](fn func(T) U, in []T) []U {
	out := make([]U, len(in))

	for i, v := range in {
		out[i] = fn(v)
	}

	return out
}
