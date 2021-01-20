package piscine

func Index(s string, toFind string) int {
	sentence := []rune(s)
	pattern := []rune(toFind)

	if len(s) < 1 {
		return 0
	}
	if len(toFind) < 1 {
		return 0
	}
	for i := range sentence {
		count := 0
		for j := range pattern {
			if sentence[i+j] != pattern[j] {
				break
			} else if sentence[i+j] == pattern[j] {
				count++
			}
			// guarantees that i is only returned if it has the same range (size, numbers of chars) of c
			//(and c is increasing over chars that are the same as i and j iterate)
			if count == len(pattern) {
				return i
				// it is returning runtime error because i (index) keeps looking for a char after the end of the string, since (count == c)
			}
		}

	}

	return -1
}
