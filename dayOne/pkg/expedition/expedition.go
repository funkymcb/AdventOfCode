package expedition

import "fmt"

type Expedition struct {
	Elfs          []Elf
	MostPackedElf Elf
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
			fmt.Println(maxCalories)
			e.SetMostPackedElf(elf)
		}
	}
}

func (e *Expedition) SetMostPackedElf(elf Elf) {
	e.MostPackedElf = elf
}
