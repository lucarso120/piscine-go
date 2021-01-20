package piscine

func TrimAtoi(s string) int {
	var a []rune
	var b []rune
	catch := []rune(s)

	for i := range catch {
		if catch[i] >= '0' && catch[i] <= '9' {
			a = append(a, catch[i])
		}
	}
	for i := range catch {
		if catch[i] >= '0' && catch[i] <= '9' || catch[i] == '-' {
			b = append(b, catch[i])
		}
	}
	intg := 0

	for i := 0; i <= len(a)-1; i++ {
		intg = (intg*10 + (int(a[i] - '0')))
	}

	for range b {
		if b[0] == '-' {
			return intg * -1

		}
	}

	return intg

}
