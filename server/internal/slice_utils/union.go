package slice_utils

// Union returns the union of two slices.
func Union[T string | int](a, b []T) []T {
	m := make(map[T]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		m[v] = struct{}{}
	}

	var result []T
	for k := range m {
		result = append(result, k)
	}
	return result
}
