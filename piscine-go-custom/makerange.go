package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil

	}
	min -= 1
	list := make([]int, max-min-1)
	for i := range list {
		list[i] = min + 1
		min += 1
	}
	return list
}
