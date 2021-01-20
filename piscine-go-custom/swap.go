package piscine

func Swap(a *int, b *int) {
	var c = *b
	*b = *a
	*a = c

}
