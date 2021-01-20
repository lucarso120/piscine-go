package piscine

func IsLower(s string) bool {
	b := []byte(s)

	for i := range s {
		if b[i] < 'a' || b[i] > 'z' {
			return false
		}
	}
	return true
}
