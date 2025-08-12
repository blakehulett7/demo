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
	Id      int
	Targets []int
}

func main() {
	fmt.Println("Dominus Iesus Christus")

	e1 := Entity{
		Id:      1,
		Targets: []int{},
	}

	e2 := Entity{
		Id:      2,
		Targets: []int{3},
	}

	e3 := Entity{
		Id:      3,
		Targets: []int{2, 5},
	}

	e4 := Entity{
		Id:      4,
		Targets: []int{23},
	}

	e5 := Entity{
		Id:      5,
		Targets: []int{3},
	}

	entities := []Entity{e1, e2, e3, e4, e5}
	for i, e := range entities {

	}
}

func GetMatchedTargets(e Entity, possible_targets []Entity) int {
	best_match := 0
	for i, t := range possible_targets {
		fmt.Printf("Checking entity %d against entity %d...", e.Id, t.Id)

		k1 := e.Id
		k2 := t.Id
		if e.Id > t.Id {
			fmt.Println("switched b/c of numerical value")
			k2 = e.Id
			k1 = t.Id
		}
		composite_key := fmt.Sprintf("%d-%d", k1, k2)

		match, exists := target_map[composite_key]
		if !exists {
			fmt.Println("no match")
			continue
		}

		fmt.Printf("match found... %d", match)

		if best_match < match {
			fmt.Printf("new best is %d", match)
			best_match = match
		}

		matched_e := Entity{Id: match}
		new_targets := possible_targets[:i]
		if i < len(possible_targets) {
			new_targets = append(possible_targets[:i], possible_targets[i+1:]...)
		}

		new_match := GetMatchedTargets(matched_e, new_targets)
		if best_match < new_match {
			best_match = new_match
		}
	}

	return best_match
}
