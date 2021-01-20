package piscine

func IsNumeric(s string) bool {
	a := []byte(s)

	for i := range a {
		if a[i] < '0' || a[i] > '9' {
			return false
		}
	}
	return true
}
