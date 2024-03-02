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

func testUrlify(f func(string, int) string) {
	var tests = []struct {
		input string
		size  int
		want  string
	}{
		{"Abhinav Robinson", 16, "Abhinav%20Robinson"},
		{"Mr John Smith       ", 13, "Mr%20John%20Smith"},
		{"     Hello    Body   ", 11, "Hello%20%20%20%20Bo"},
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
}
