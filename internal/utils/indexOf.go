package utils

func IndexOf[T comparable](arr []T, value T, startPoint int) int {
	for index := startPoint; index < len(arr); index++ {
		if arr[index] == value {
			return index
		}
	}
	return -1
}
