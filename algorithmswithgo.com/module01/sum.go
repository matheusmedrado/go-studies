package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	var sum int // we could also use sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

// // There's also a recurssion method that we could use

// func Sum(numbers []int) int {
// 	if len(numbers) == 0{
// 		return 0
// 	}
	
// 	return numbers[0] + Sum(numbers[1:])

// }