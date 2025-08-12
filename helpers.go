package main

func GenerateArrayPermutations(input []int) [][]int {
	if len(input) != 5 {
		panic("should only ever have 5 entity id's")
	}

	var result [][]int
	n := len(input)

	// Recursive helper function
	var generate func(k int, currentPermutation []int)
	generate = func(k int, currentPermutation []int) {
		if k == 1 {
			// Base case: a permutation is found, make a copy and add to result
			temp := make([]int, n)
			copy(temp, currentPermutation)
			result = append(result, temp)
			return
		}

		// Recursively generate permutations
		for i := 0; i < k; i++ {
			generate(k-1, currentPermutation)

			// Swap elements to generate next permutation
			if k%2 == 1 { // Odd length
				currentPermutation[0], currentPermutation[k-1] = currentPermutation[k-1], currentPermutation[0]
			} else { // Even length
				currentPermutation[i], currentPermutation[k-1] = currentPermutation[k-1], currentPermutation[i]
			}
		}
	}

	// Start the generation with a copy of the original slice
	initialPermutation := make([]int, n)
	copy(initialPermutation, input)
	generate(n, initialPermutation)

	if len(result) != 120 {
		panic("did not find all 120 permutations")
	}

	return result
}
