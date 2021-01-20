package piscine

//alternative version
//if index < {
//	return 0
//}
//if index == 1 {
//	return 1
//}
func Fibonacci(index int) int {

	if index < 0 {
		return -1
	}
	if index < 2 {
		return index
	}

	return Fibonacci(index-1) + Fibonacci(index-2)

}
