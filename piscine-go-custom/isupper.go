package piscine

func IsUpper(s string) bool {
	a := []byte(s)
	for i := range s {
		if a[i] < 'A' || a[i] > 'Z' {
			return false
		}

	}
	return true
}
