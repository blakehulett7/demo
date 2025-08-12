package main

import (
	"fmt"
	"slices"
	"sort"
)

var target_map map[string]int = map[string]int{
	"2-3":   23,
	"3-5":   35,
	"9-9":   99,
	"4-23":  423,
	"15-64": 1564,
	"23-44": 2344,
	"65-99": 6599,
}
var possible_values = []int{0, 23, 35, 99, 423, 1564, 2344, 6599}

const HighestValue = 6599

func main() {
	fmt.Println("Dominus Iesus Christus")

	entities := []int{1, 5, 9, 9, 65}
	best_match := FindBestMatch(entities)

	fmt.Println(best_match)
}

func FindBestMatch(entities []int) int {
	sort.Ints(entities)

	best_match := 0
	for i, e := range entities {
		targets := make([]int, 5)
		copy(targets, entities)
		targets = slices.Delete(targets, 0, i+1)

		match := FindMatches(e, targets)
		if best_match < match {
			best_match = match
		}
	}

	return best_match
}

func FindMatches(entity int, targets []int) int {
	if len(targets) == 0 {
		return 0
	}

	best_match := 0
	for i, target := range targets {
		k1 := entity
		k2 := target

		if k2 < k1 {
			k1 = target
			k2 = entity
		}

		composite_key := fmt.Sprintf("%d-%d", k1, k2)
		match, found := target_map[composite_key]
		if !found {
			continue
		}

		if best_match < match {
			best_match = match
		}

		nested_targets := make([]int, len(targets))
		copy(nested_targets, targets)
		nested_targets = slices.Delete(nested_targets, i, i+1)
		match = FindMatches(match, nested_targets)

		if best_match < match {
			best_match = match
		}
	}

	return best_match
}
