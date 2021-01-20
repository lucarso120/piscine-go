package piscine

func ToUpper(s string) string {
	a := []rune(s)

	for i := range s {
		if a[i] >= 'a' && a[i] <= 'z' {
			a[i] = a[i] - 32
		}

	}
	b := string(a)
	return b
}
