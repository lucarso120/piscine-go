package piscine

func Capitalize(s string) string {
	a := []rune(s)

	for i := range s {
		if a[i] >= 'A' && a[i] <= 'Z' {
			a[i] = a[i] + 32
		}
		if a[0] >= 'a' && a[0] <= 'z' {
			a[i] = a[i] - 32
		}
		if (a[i] >= 'a' && a[i] <= 'z') && (a[i-1] < 48) {
			a[i] = a[i] - 32
		} else if (a[i] >= 'a' && a[i] <= 'z') && (a[i-1] > 122) {
			a[i] = a[i] - 32
		} else if (a[i] >= 'a' && a[i] <= 'z') && (a[i-1] > 57 && a[i-1] < 65) {
			a[i] = a[i] - 32
		} else if (a[i] >= 'a' && a[i] <= 'z') && (a[i-1] > 90 && a[i-1] < 97) {
			a[i] = a[i] - 32
		}
	}
	b := string(a)
	return b

}
