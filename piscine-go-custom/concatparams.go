package piscine

func ConcatParams(args []string) string {
	strings := ""

	count := 0
	for range args {
		count++
	}

	for i, str := range args {
		strings += string(str)
		if i < count-1 {
			strings += "\n"

		}

	}
	return strings

}
