package main

// Sum returns the sum of numbers in the array
func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}