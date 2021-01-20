package piscine

import (
	"fmt"
	"os"
)

func main() {
	argument := os.Args[1:]

	list := []string{}

	for i := range argument {
		list = append(list, argument[i])
	}

	for i := range list {
		fmt.Println(list[i])

	}
}
