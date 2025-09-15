package main

import "testing"

var testMeanCalculatorArrayInputs = []struct {
	input    [5]int
	expected int
}{
	{[5]int{1, 2, 3, 4, 5}, 3},
	{[5]int{10, 20, 30, 40, 50}, 30},
	{[5]int{100, 100, 100, 100, 100}, 100},
	{[5]int{1000, 2000, 3000, 4000, 5000}, 3000},
}
var testMeanCalculatorSliceInputs = []struct {
	input    []int
	expected int
}{
	{[]int{1, 2, 3, 4, 5}, 3},
	{[]int{10, 20, 30, 40, 50}, 30},
	{[]int{100, 100, 100, 100, 100}, 100},
	{[]int{1000, 2000, 3000, 4000, 5000}, 3000},
}

func TestMeanCalculatorArrayInputs(t *testing.T) {
	for _, value := range testMeanCalculatorArrayInputs {
		result := meanCalculatorArray(value.input)
		expected := value.expected
		if result != expected {
			t.Errorf("Error Calc = %d; want %d", result, expected)
		}
	}
}
func TestMeanCalculatorSliceInputs(t *testing.T) {
	for _, value := range testMeanCalculatorSliceInputs {
		result := meanCalculatorSlice(value.input)
		expected := value.expected
		if result != expected {
			t.Errorf("Error Calc = %d; want %d", result, expected)
		}
	}
}
func BenchmarkMeanCalculatorArray(b *testing.B) {
	for _, value := range testMeanCalculatorArrayInputs {
		for i := 0; i < b.N; i++ {
			_ = meanCalculatorArray(value.input)
		}
	}
}
func BenchmarkMeanCalculatorSlice(b *testing.B) {
	for _, value := range testMeanCalculatorSliceInputs {
		for i := 0; i < b.N; i++ {
			_ = meanCalculatorSlice(value.input)
		}
	}
}
func BenchmarkForArray(b *testing.B) {
	for _, value := range testMeanCalculatorArrayInputs {
		for i := 0; i < b.N; i++ {
			_ = usingForArray(value.input)
		}
	}
}
func BenchmarkForSlice(b *testing.B) {
	for _, value := range testMeanCalculatorSliceInputs {
		for i := 0; i < b.N; i++ {
			_ = usingForSlice(value.input)
		}
	}
}

func BenchmarkForRangeArray(b *testing.B) {
	for _, value := range testMeanCalculatorArrayInputs {
		for i := 0; i < b.N; i++ {
			_ = usingForRangeArray(value.input)
		}
	}
}
func BenchmarkForRangeSlice(b *testing.B) {
	for _, value := range testMeanCalculatorSliceInputs {
		for i := 0; i < b.N; i++ {
			_ = usingForRangeSlice(value.input)
		}
	}
}
