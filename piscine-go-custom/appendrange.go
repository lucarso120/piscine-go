package piscine

func AppendRange(min, max int) []int {
	var list []int
	for i := min; i < max; i++ {
		list = append(list, i)
	}
	return list
}
