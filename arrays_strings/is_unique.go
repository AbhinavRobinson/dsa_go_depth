package main

import "fmt"

func isUnique(input string) bool {
	memory := make(map[int32]bool)
	for _, ascii := range input {
		if memory[ascii] {
			return false
		}
		memory[ascii] = true
	}
	return true
}

func isUniqueNoExternalDataStructure(input string) bool {
	for index := range input {
		for _, compare := range input[:index] {
			if input[compare] == input[index] {
				return false
			}
		}
	}
	return true
}

func testUnique(f func(string) bool) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"abhinav", false},
		{"bingo", true},
	}
	for _, tt := range tests {
		ans := f(tt.input)
		if ans != tt.want {
			fmt.Printf("Failed %s\n", tt.input)
		} else {
			fmt.Printf("Passed %s\n", tt.input)
		}
	}
}

func main() {
	fmt.Println("Is unique:")
	testUnique(isUnique)
	fmt.Println("\nIs unique with no external ds:")
	testUnique(isUniqueNoExternalDataStructure)
}
