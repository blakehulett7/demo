package main

import (
	"slices"
	"testing"
)

func FuzzBestMatch(f *testing.F) {
	seed_tests := [][]int{
		{1, 2, 3, 4, 5},
		{1, 6, 5, 9, 9},
		{1, 4, 5, 6, 7},
		{1, 3, 5, 9, 9},
	}

	for _, seed_test := range seed_tests {
		f.Add(seed_test[0], seed_test[1], seed_test[2], seed_test[3], seed_test[4])
	}
	f.Fuzz(func(t *testing.T, a int, b int, c int, d int, e int) {
		input := []int{a, b, c, d, e}
		best_match := FindBestMatch(input)

		if best_match > HighestValue {
			t.Errorf("Got value: %d that exceed highest possible value\n", best_match)
		}

		if !slices.Contains(possible_values, best_match) {
			t.Errorf("Got value: %d which is not on the possible values list\n", best_match)
		}
	})
}
