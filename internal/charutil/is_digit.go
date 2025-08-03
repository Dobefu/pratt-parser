package charutil

// IsDigit checks if a rune is a digit.
func IsDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
