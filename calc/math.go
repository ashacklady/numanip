package calc

import "errors"

// func Add(i, j int) int {
// 	return i + j
// }

// func Add(numbers ...int) int {
// 	sum := 0
// 	for _, num := range numbers {
// 		sum += num
// 	}
// 	return sum
// }

func Add(numbers ...int) (error, int) {
	sum := 0
	if len(numbers) < 2 {
		return errors.New("provide more than 2 numbers"), sum
	}
	for _, num := range numbers {
		sum += num
	}
	return nil, sum
}
