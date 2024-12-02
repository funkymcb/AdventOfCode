package day2

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/funkymcb/AdventOfCode/internal/io"
)

const DAY = 2

type Day2 struct{}

func handleInput() ([][]int, error) {
	var reports [][]int

	lines, err := io.ReadFile(DAY)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {

		levels := strings.Fields(line)

		levelsInt := make([]int, 0, len(levels))
		for _, level := range levels {
			num, err := strconv.Atoi(level)
			if err != nil {
				return nil, fmt.Errorf("error converting %s to int: %w", line, err)
			}
			levelsInt = append(levelsInt, num)
		}

		reports = append(reports, levelsInt)
	}

	return reports, nil
}

func analyzeReports(reports [][]int) (int, error) {
	var result int

Reports:
	for _, report := range reports {
		var gaps []int
		for i, level := range report {
			if i == 0 {
				//first level can be skipped
				continue
			}

			gaps = append(gaps, (level - report[i-1]))
		}

		allIncreasing := true
		allDecreasing := true

		for _, gap := range gaps {
			if gap == 0 || gap > 3 || gap < -3 {
				continue Reports
			}
			if gap > 0 {
				allDecreasing = false
			}
			if gap < 0 {
				allIncreasing = false
			}
		}

		if (allIncreasing && !allDecreasing) || (!allIncreasing && allDecreasing) {
			result++
		}
	}

	return result, nil
}

func (d Day2) Run() {
	reports, err := handleInput()

	star1, err := analyzeReports(reports)
	if err != nil {
		fmt.Println("error analyzing reactor reports", err)
		os.Exit(1)
	}

	var star2 int

	io.PrintResult(reflect.TypeOf((*Day2)(nil)).Elem().Name(), star1, star2)
}
