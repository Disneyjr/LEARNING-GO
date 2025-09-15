package main

import "fmt"

func main() {
	ContactsCRUD := make(map[string]string)
	createContacts("Teacher Go", "+55679123456", ContactsCRUD)
	if len(ContactsCRUD) > 0 {
		readContacts("Teacher Go", ContactsCRUD)
		updateContacts("Teacher Go", "+55999999999", ContactsCRUD)
		readContacts("Teacher Go", ContactsCRUD)
		deleteContacts("Teacher Go", ContactsCRUD)
		readContacts("Teacher Go", ContactsCRUD)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Struct Person")
	firstPerson := Person{"Lucas", 12}

	fmt.Println(firstPerson.name)
	fmt.Println(firstPerson.age)
	Birthday(&firstPerson)
	fmt.Println(getPrint(firstPerson))

}
func meanCalculatorArray(numbers [5]int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum / len(numbers)
}
func meanCalculatorSlice(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum / len(numbers)
}
