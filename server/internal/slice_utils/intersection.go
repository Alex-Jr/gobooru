package slice_utils

// Intersection returns the intersection of two slices.
func Intersection[T string | int](a, b []T) []T {
	m := make(map[T]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}

	var result []T
	for _, v := range b {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
