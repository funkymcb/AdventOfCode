package day2

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/funkymcb/AdventOfCode/internal/io"
)

func handleInput(day int) ([][]int, error) {
	var reports [][]int

	lines, err := io.ReadFile(day)
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

func reportIsSafe(report []int) bool {
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
	isOK := true

	for _, gap := range gaps {
		if gap == 0 || gap > 3 || gap < -3 {
			isOK = false
			continue
		}
		if gap > 0 {
			allDecreasing = false
		}
		if gap < 0 {
			allIncreasing = false
		}
	}
	if isOK && (allIncreasing || allDecreasing) {
		return true
	}
	return false
}

func analyzeReports(reports [][]int) int {
	var safeReports int

	for _, report := range reports {
		if reportIsSafe(report) {
			safeReports++
		}
	}

	return safeReports
}

func removeLevel(report []int, index int) []int {
	return append(report[:index], report[index+1:]...)
}

func generateDampedReports(report []int) [][]int {
	var dampedReports [][]int

	for i, _ := range report {
		dr := slices.Clone(report)
		dampedReports = append(dampedReports, removeLevel(dr, i))
	}

	return dampedReports
}

func applyProblemDampener(reports [][]int, safeReportCount int) int {
	var safeReports int

	for _, report := range reports {
		if reportIsSafe(report) {
			safeReports++
		} else {
			// remove faulty element
			dampedReports := generateDampedReports(report)
			for _, dampedReport := range dampedReports {
				if reportIsSafe(dampedReport) {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports
}

func Run(day int) {
	reports, err := handleInput(day)
	if err != nil {
		fmt.Println("error handling input", err)
	}

	star1 := analyzeReports(reports)

	star2 := applyProblemDampener(reports, star1)

	io.PrintResult(day, star1, star2)
}
