package piscine

func AlphaCount(s string) int {
	c := 0

	for _, i := range s {
		if (i >= 'A' && i <= 'Z') || (i >= 'a' && i <= 'z') {
			c++
		}
	}
	return c
}
