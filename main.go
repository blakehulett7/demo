package main

import (
	"fmt"
)

var target_map map[string]int = map[string]int{
	"2-3":  23,
	"3-5":  35,
	"4-23": 423,
}

type Entity struct {
	Id int
}

func main() {
	fmt.Println("Dominus Iesus Christus")

	entities := []int{1, 2, 3, 4, 5}
	best_match := GetBestMatch(entities)

	fmt.Println(best_match)
}

func GetBestMatch(targets []int) int {
	best_match := 0
	if len(targets) == 0 {
		return 0
	}

	current := targets[0]

	for i := 1; i < len(targets); i++ {
		target := targets[i]

		k1 := current
		k2 := target
		if target < current {
			k1 = target
			k2 = current
		}

		composite_key := fmt.Sprintf("%d-%d", k1, k2)
		fmt.Printf("Checking composite key: %s\n", composite_key)

		match, exists := target_map[composite_key]
		if !exists {
			fmt.Println("no match found...")
			continue
		}

		fmt.Printf("found match: %d\n", match)
		if best_match < match {
			best_match = match
		}
	}

}

func GetMatchedTargets(current_i int, e Entity, possible_targets []Entity) int {
	best_match := 0
	fmt.Printf("Checking entity %d against possible targets %v\n", e, possible_targets)
	for i, t := range possible_targets {
		if i <= current_i {
			continue
		}

		fmt.Printf("Checking entity %d against entity %d...\n", e.Id, t.Id)

		targets_copy := make([]Entity, len(possible_targets))
		copy(targets_copy, possible_targets)

		k1 := e.Id
		k2 := t.Id
		if e.Id > t.Id {
			fmt.Println("switched b/c of numerical value")
			k2 = e.Id
			k1 = t.Id
		}
		composite_key := fmt.Sprintf("%d-%d", k1, k2)
		fmt.Printf("Composite Key: %s\n", composite_key)

		match, exists := target_map[composite_key]
		if !exists {
			fmt.Println("no match")
			continue
		}

		fmt.Printf("match found... %d\n", match)

		if best_match < match {
			fmt.Printf("new best is %d\n", match)
			best_match = match
		}

		matched_e := Entity{Id: match}

		new_targets := []Entity{}
		for j, new_t := range possible_targets {
			if j == i || j == current_i {
				continue
			}

			new_targets = append(new_targets, new_t)
		}

		new_match := GetMatchedTargets(0, matched_e, new_targets)
		if best_match < new_match {
			best_match = new_match
		}
	}

	return best_match
}
