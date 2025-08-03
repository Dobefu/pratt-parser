package charutil

// IsLetter checks if a rune is a letter.
func IsLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}
