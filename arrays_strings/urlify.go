package main

import (
	"fmt"
)

func urlify(input string, size int) string {
	result := ""
	count := 0
	for index := 0; index < len(input) && count < size; index++ {
		if len(result) == 0 && input[index] == ' ' {
			continue
		}
		if input[index] == ' ' {
			result += "%20"
		} else {
			result += string(input[index])
		}
		count++
	}
	return result
}

func urlifyInplace(input string, size int) string {
	spaces := 0
	bInput := []byte(input)
	for index := 0; index < size; index++ {
		if input[index] == ' ' {
			spaces++
		}
	}
	index := size + spaces*2
	for position := size - 1; position >= 0; position-- {
		if bInput[position] == ' ' {
			bInput[index-3] = '%'
			bInput[index-2] = '2'
			bInput[index-1] = '0'
			index -= 3
		} else {
			bInput[index-1] = bInput[position]
			index--
		}
	}
	return string(bInput)
}

func testUrlify(f func(string, int) string) {
	var tests = []struct {
		input string
		size  int
		want  string
	}{
		{"Abhinav Robinson  ", 16, "Abhinav%20Robinson"},
		{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
	}
	for _, tt := range tests {
		ans := f(tt.input, tt.size)
		if ans != tt.want {
			fmt.Printf("Failed %s : %s <- %s\n", tt.input, ans, tt.want)
		} else {
			fmt.Printf("Passed %s\n", tt.input)
		}
	}
}

func main() {
	fmt.Println("URLify")
	testUrlify(urlify)
	fmt.Println("\nURLify inplace")
	testUrlify(urlifyInplace)
}
