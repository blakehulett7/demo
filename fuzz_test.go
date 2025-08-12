package main

import (
	"testing"
)

func FuzzBestMatch(f *testing.F) {
	seed_tests := [][]int{
		{1, 2, 3, 4, 5},
	}

	f.Add(seed_tests[0][0], seed_tests[0][1], seed_tests[0][2], seed_tests[0][3], seed_tests[0][4])
	f.Fuzz(func(t *testing.T, a int, b int, c int, d int, e int) {
		input := []int{a, b, c, d, e}
		permutations := GenerateArrayPermutations(input)

		best_match := FindBestMatch(permutations[0])
		for _, tc := range permutations {
			current_best := FindBestMatch(tc)

			if current_best != best_match {
				t.Errorf("All permutations should return the same answer. Baseline best: %d, Failing permutation: %v\n", best_match, tc)
			}
		}
	})
}
