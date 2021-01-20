package piscine

func LastRune(s string) rune {
	a := []rune(s)
	b := len(s)

	return (a[b-1])
}
