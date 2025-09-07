package main

import (
	"errors"
	"fmt"
)

func main() {
	FirstExercise()
	fmt.Print("\n")
	SecondExercise()
	fmt.Print("\n")
	result1, result2 := ThirdExercise(5, 8)
	fmt.Printf("a + b %d, %d", result1, result2)
	resultFourExercise, err := FourthExercise(6, 2)
	fmt.Print("\n")
	if err != nil {
		fmt.Printf("Result of the division: %s\n", err)
	} else {
		fmt.Printf("Result of the division: %d\n", resultFourExercise)
	}

	resultFourExerciseErr, err := FourthExercise(6, 0)
	if err != nil {
		fmt.Printf("Result of the division: %s\n", err)
	} else {
		fmt.Printf("Result of the division: %d\n", resultFourExerciseErr)
	}
}

func FirstExercise() {
	nome := "Disney Junior"
	fmt.Printf("Hello, <%s>! Wellcome to Go.", nome)
}

func SecondExercise() {
	firstValue := 10
	secondValue := 20
	thirdValue := 50
	result := 0
	result = firstValue + secondValue
	fmt.Printf("10 + 20: %d\n", result)
	result = thirdValue - result
	fmt.Printf("50 - 30: %d\n", result)
	fmt.Printf("Resultado: %d", result)
}

func ThirdExercise(a, b int) (int, int) {
	firstValue := a
	a = b
	b = firstValue
	return a, b
}
func FourthExercise(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
