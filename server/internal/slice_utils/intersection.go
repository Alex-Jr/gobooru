package slice_utils

// A ∩ B, returns an slice with only elements that are in both A and B
func Intersection[T string | int](a, b []T) []T {
	m := make(map[T]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}

	result := make([]T, 0)
	for _, v := range b {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
