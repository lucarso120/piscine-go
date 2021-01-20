package piscine

func Compare(a, b string) int {
	x, y := []byte(a), []byte(b)
	for i := range x {
		for j := range y {
			if a == b {
				return 0
			} else if x[i] < y[j] {
				return -1
			} else {
				return 1
			}
		}

	}
	return 0
}
