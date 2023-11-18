package slice_utils

// A  B, returns an slice with the elements that are in A but not in B
func Difference[T string | int](a, b []T) []T {
	m := make(map[T]bool)
	for _, item := range b {
		m[item] = true
	}

	diff := make([]T, 0)
	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}

	return diff
}
