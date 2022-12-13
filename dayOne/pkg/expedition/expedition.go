package expedition

import (
	"sort"
)

type Expedition struct {
	Elves                  []Elf
	MostPackedElf          Elf
	CaloriesOfTopThreeElfs int
}

type Elf struct {
	ID            int
	Meals         []int
	TotalCalories int
}

func (e *Expedition) GetMostPackedElf() {
	var maxCalories int
	for _, elf := range e.Elves {
		if elf.TotalCalories > maxCalories {
			maxCalories = elf.TotalCalories
			e.SetMostPackedElf(elf)
		}
	}
}

func (e *Expedition) SetMostPackedElf(elf Elf) {
	e.MostPackedElf = elf
}

func (e *Expedition) SortElfsByTotalCalories() {
	sort.Slice(e.Elves, func(i, j int) bool {
		return e.Elves[i].TotalCalories > e.Elves[j].TotalCalories
	})
}

func (e *Expedition) GetCaloriesOfTopThreeElfs() {
	var total int
	for i := 0; i < 3; i++ {
		total = total + e.Elves[i].TotalCalories
	}
	e.CaloriesOfTopThreeElfs = total
}
