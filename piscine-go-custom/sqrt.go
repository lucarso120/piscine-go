package piscine

func Sqrt(nb int) int {
	number := 0
	if nb == 1 {
		return 1
	}
	if nb <= 100 {
		for i := 0; i <= nb-1; i++ {
			for j := 0; j <= nb-1; j++ {
				if j == i && j*i == nb {
					number = i
				}
			}
		}

	}

	return number
}
