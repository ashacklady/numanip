package calc

// func Add(i, j int) int {
// 	return i + j
// }

func Add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
