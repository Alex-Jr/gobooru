package slice_utils

// TwoWayDifference returns the difference between two slices.
// It returns the elements that are in a but not in b and the elements that are in b but not in a.
func TwoWayDifference[T string | int](a, b []T) ([]T, []T) {
	inA := make([]T, 0)
	inB := make([]T, 0)

	aMap := make(map[T]struct{}, len(a))
	bMap := make(map[T]struct{}, len(b))

	for _, aVal := range a {
		aMap[aVal] = struct{}{}
	}

	for _, bVal := range b {
		bMap[bVal] = struct{}{}
	}

	for aVal := range aMap {
		if _, ok := bMap[aVal]; !ok {
			inA = append(inA, aVal)
		}
	}

	for bVal := range bMap {
		if _, ok := aMap[bVal]; !ok {
			inB = append(inB, bVal)
		}
	}

	return inA, inB
}
