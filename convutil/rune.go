package conv

// Rune converts `any` to rune.
func Rune(any any) (rune, error) {
	if v, ok := any.(rune); ok {
		return v, nil
	}
	return Int32(any)
}

// Runes converts `any` to []rune.
func Runes(any any) []rune {
	if v, ok := any.([]rune); ok {
		return v
	}
	return []rune(String(any))
}
