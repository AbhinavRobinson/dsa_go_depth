package main

import "fmt"

// V.E12 A function that calculates all possible permutations of a string.

func permutation(str string, prefix string) { // O(N^2 * N!)
	if len(str) == 0 {
		fmt.Println(prefix) // O(N)
	} else {
		for i := 0; i < len(str); i++ { // O(N)
			rem := str[0:i] + str[i+1:]             // O(1)
			permutation(rem, prefix+string(str[i])) // N! Roots, each cost N -> O(N * N!)
		}
	}
}

func permutations(str string) {
	permutation(str, "")
}

func main() {
	permutations("test")
}
