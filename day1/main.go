package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/funkymcb/AdventOfCode/day1/pkg/expedition"
)

var tmpElf expedition.Elf
var exp expedition.Expedition

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func sumOfTotalCalories(m []int) int {
	var total int
	for _, meal := range m {
		total = total + meal
	}
	return total
}

func appendElf(id int, meals []int) {
	totalCalories := sumOfTotalCalories(meals)

	exp.Elves = append(exp.Elves, expedition.Elf{
		ID:            id,
		Meals:         meals,
		TotalCalories: totalCalories,
	})
}

func scanInput(input string) {
	var id = 1
	var meals = []int{}

	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		if line == "" {
			appendElf(id, meals)
			id++
			meals = []int{}
			continue
		}

		meal, err := strconv.Atoi(line)
		checkError(err)

		meals = append(meals, meal)
	}
	appendElf(id, meals)
}

func printExpedition() {
	for _, elf := range exp.Elves {
		fmt.Printf("Elf #%d:\nMeals:%v\nTotal Calories:%d\n\n", elf.ID, elf.Meals, elf.TotalCalories)
	}
	fmt.Printf("Most packed Elf: Elf #%d with %d Calories\n\n", exp.MostPackedElf.ID, exp.MostPackedElf.TotalCalories)

	fmt.Printf("Calories of top three Elfs: %d\n", exp.CaloriesOfTopThreeElfs)
}

func main() {
	input, err := os.ReadFile("./list_of_supplies")
	checkError(err)

	scanInput(string(input))

	exp.GetMostPackedElf()
	exp.SortElfsByTotalCalories()
	exp.GetCaloriesOfTopThreeElfs()
	printExpedition()
}
