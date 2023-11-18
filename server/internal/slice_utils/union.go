package slice_utils

// A + B, returns an slice with elements from A and B, will deduplicate amy.
func Union[T string | int](a, b []T) []T {
	m := make(map[T]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		m[v] = struct{}{}
	}

	result := make([]T, 0)
	for k := range m {
		result = append(result, k)
	}
	return result
}
