package slice_utils

func Deduplicate[T string | int](a []T) []T {
	m := make(map[T]bool)
	for _, item := range a {
		m[item] = true
	}

	unique := make([]T, len(m))
	i := 0
	for item := range m {
		unique[i] = item
		i++
	}

	return unique
}
