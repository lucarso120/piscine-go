package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var result []int

	if n == 0 {
		z01.PrintRune(rune('0'))
	}

	for n > 0 {
		a := n % 10
		n = n / 10
		result = append(result, a)
	}
	//loop in ascending order
	//adds the value result[i] to the 48 on ascii table. The result, is, thus, a rune
	for j := 0; j <= 9; j++ {
		for i := range result {
			if result[i] == j {
				z01.PrintRune(rune('0' + result[i]))

			}
		}
	}

}
