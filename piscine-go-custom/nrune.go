package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	b := 0
	for range s {
		b++
		if b == n {
			return a[b-1]
		}
	}
	return 0

}
