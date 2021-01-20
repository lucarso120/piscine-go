package piscine

func BasicAtoi(s string) int {
	a := []byte(s)
	n := len(s)
	intg := 0

	for i := 0; i <= n-1; i++ {
		intg = intg*10 + (int(a[i] - '0'))

	}
	return intg
}
