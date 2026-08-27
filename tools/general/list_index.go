package general

// IndexOf to get index of list
func IndexOf[T comparable](data []T, element T) int {
	for i, v := range data {
		if v == element {
			return i
		}
	}
	return -1
}
