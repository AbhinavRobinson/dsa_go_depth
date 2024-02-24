package main

import "fmt"

func checkPermutation(source string, compare string) bool {
	if len(source) != len(compare) {
		return false
	}
	for index := range source {
		found := false
		for target := range compare {
			if compare[target] == source[index] {
				found = true
				compare = compare[0:target] + compare[target+1:]
				break
			}
		}
		if !found {
			return false
		}
	}
	if len(compare) > 0 {
		return false
	}
	return true
}

func testPermutation(f func(string, string) bool) {
	var tests = []struct {
		input  string
		target string
		want   bool
	}{
		{"church", "hccuhr", true},
		{"golang", "goland", false},
		{"somelargewordthatshouldbeapermutation", "tmemutatordshouhatsoldbeaplargewerion", true},
		{"somelargewordthatshouldnotbeapermutation", "somelargewordthatshoulddotbeapermutation", false},
		{"somelargewordthatshouldnotbeapermutationshorter", "somelargewordthatshouldnotbeapermutationlongerword", false},
	}
	for _, tt := range tests {
		ans := f(tt.input, tt.target)
		if ans != tt.want {
			fmt.Printf("Failed %s\n", tt.input)
		} else {
			fmt.Printf("Passed %s\n", tt.input)
		}
	}
}

func main() {
	testPermutation(checkPermutation)
}
