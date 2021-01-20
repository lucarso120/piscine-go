package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	for i := nb; i > 0; i++ {
		for j := 2; i > 0; j++ {
			if i != j && i%j == 0 {
				break
			} else if i == j {
				return i
			}

		}
	}
	return 0

}
