package piscine

func IterativePower(nb int, power int) int {
	parameter := nb

	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1

	}

	for i := 1; i < power; i++ {
		parameter = parameter * nb
	}

	return parameter

}
