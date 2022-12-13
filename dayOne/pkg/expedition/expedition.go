package expedition

import (
	"sort"
)

type Expedition struct {
	Elfs                   []Elf
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
	for _, elf := range e.Elfs {
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
	sort.Slice(e.Elfs, func(i, j int) bool {
		return e.Elfs[i].TotalCalories > e.Elfs[j].TotalCalories
	})
}

func (e *Expedition) GetCaloriesOfTopThreeElfs() {
	var total int
	for i := 0; i < 3; i++ {
		total = total + e.Elfs[i].TotalCalories
	}
	e.CaloriesOfTopThreeElfs = total
}
