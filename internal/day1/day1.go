package day1

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/funkymcb/AdventOfCode/internal/io"
)

const DAY = 1

type Day1 struct {
	name int
}

func handleInput() ([]int, []int, error) {
	var locationIDleft, locationIDright []int

	lines, err := io.ReadFile(DAY)
	if err != nil {
		return nil, nil, err
	}

	for _, line := range lines {
		IDs := strings.Fields(line)

		if len(IDs) != 2 {
			fmt.Println("Line has more than 2 IDs. Skipping...", line)
			continue
		}

		IDleft, err1 := strconv.Atoi(IDs[0])
		IDright, err2 := strconv.Atoi(IDs[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("error converting input ID to int, %s, %s", err1, err2)
		}

		locationIDleft = append(locationIDleft, IDleft)
		locationIDright = append(locationIDright, IDright)
	}

	// if err := file.Err(); err != nil {
	// 	return nil, nil, fmt.Errorf("error processing lines of input file: %s", err)
	// }

	return locationIDleft, locationIDright, nil
}

// day 1 star 1
func calculateTotalDistances(IDsLeft, IDsRight []int) (int, error) {
	var result int

	sort.Ints(IDsLeft)
	sort.Ints(IDsRight)

	for i, ID := range IDsLeft {
		distance := ID - IDsRight[i]

		if distance < 0 {
			distance *= -1
		}

		result += distance
	}

	return result, nil
}

// day 1 star 2
func calculateSimilarityScore(IDsLeft, IDsRight []int) (int, error) {
	var result int

	for _, il := range IDsLeft {
		var factor int

		for _, ir := range IDsRight {
			if il == ir {
				factor++
			}
		}

		result += (il * factor)
	}

	return result, nil
}

func (d Day1) Run() {
	locationIDsleft, locationIDright, err := handleInput()
	if err != nil {
		fmt.Println("error interpreting input:", err)
		os.Exit(1)
	}

	star1, err := calculateTotalDistances(locationIDsleft, locationIDright)
	if err != nil {
		fmt.Println("error calculating distance of IDs", err)
		os.Exit(1)
	}

	star2, err := calculateSimilarityScore(locationIDsleft, locationIDright)
	if err != nil {
		fmt.Println("error calculating similarity score", err)
		os.Exit(1)
	}

	io.PrintResult(reflect.TypeOf((*Day1)(nil)).Elem().Name(), star1, star2)
}
