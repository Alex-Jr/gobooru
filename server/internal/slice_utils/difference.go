package slice_utils

// Diference returns the difference between two slices.
func Difference[T string | int](a, b []T) []T {
	m := make(map[T]bool)
	for _, item := range a {
		m[item] = true
	}

	var diff []T
	for _, item := range b {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}

	return diff
}
