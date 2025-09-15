package main

func usingForArray(numbers [5]int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
func usingForSlice(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
func usingForRangeArray(numbers [5]int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func usingForRangeSlice(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
