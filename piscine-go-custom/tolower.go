package piscine

func ToLower(s string) string {
	a := []byte(s)

	for i := range s {
		if a[i] >= 'A' && a[i] <= 'Z' {
			a[i] = a[i] + 32
		}
	}
	b := string(a)
	return b

}
