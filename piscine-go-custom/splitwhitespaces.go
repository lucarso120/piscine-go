package piscine

func SplitWhiteSpaces(s string) []string {

	var sliceStr []string

	count := 0

	for i, j := range s {

		if string(j) == " " || string(j) == "\n" || string(j) == "\t" {
			if i != 0 && s[i-1] != ' ' && s[i-1] != '\n' && s[i-1] != '\t' {

				sliceStr = append(sliceStr, s[count:i])
				count = i + 1

			}

		}

	}
	sliceStr = append(sliceStr, s[count:])

	return sliceStr

}
