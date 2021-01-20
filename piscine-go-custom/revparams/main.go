package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	lenghtArgs := len(arguments)

	for i := lenghtArgs - 1; i > 0; i-- {
		parameter := arguments[i]

		for _, prm := range parameter {
			z01.PrintRune(prm)
			z01.PrintRune('\n')
		}

	}

}
