package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := []rune(os.Args[0])

	for i := range a {
		z01.PrintRune(a[i])
	}
	z01.PrintRune('\n')
}
