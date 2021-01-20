package piscine

func IsPrintable(s string) bool {
	a := []byte(s)

	for i := range s {
		if a[i] < 32 {
			return false
		}
	}
	return true
}
