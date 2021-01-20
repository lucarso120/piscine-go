package piscine

func IsAlpha(s string) bool {
	a := []byte(s)

	for i := range a {
		if a[i] < 48 || a[i] > 122 {
			return false
		} else if a[i] >= 58 && a[i] <= 64 {
			return false
		} else if a[i] >= 91 && a[i] <= 96 {
			return false
		}
	}
	return true
}
