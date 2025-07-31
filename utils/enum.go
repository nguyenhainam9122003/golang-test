package utils

func IsValidEnum[T comparable](value T, validValues []T) bool {
	for _, v := range validValues {
		if value == v {
			return true
		}
	}
	return false
}
