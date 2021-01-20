package piscine

func StrRev(s string) string {
	a := []rune(s)
	n := 0
	for range s {
		n++
	}
	for x, y := 0, n-1; x < y; x, y = x+1, y-1 {
		a[x], a[y] = a[y], a[x]

	}

	return string(a)

}
