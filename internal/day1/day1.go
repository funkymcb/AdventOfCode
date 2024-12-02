package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() ([]int, []int) {
	file, err := os.Open("./inputs/day_1.txt")
	if err != nil {
		log.Fatalln("error opening file", err)
	}
	defer file.Close()

	var locationIDleft, locationIDright []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		IDs := strings.Fields(line)

		if len(IDs) != 2 {
			log.Println("Line has more than 2 IDs. Skipping...", line)
			continue
		}

		IDleft, err1 := strconv.Atoi(IDs[0])
		IDright, err2 := strconv.Atoi(IDs[1])
		if err1 != nil || err2 != nil {
			log.Fatalln("error converting input ID to int", err1, err2)
		}

		locationIDleft = append(locationIDleft, IDleft)
		locationIDright = append(locationIDright, IDright)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("error processing lines of input file", err)
	}

	return locationIDleft, locationIDright
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

func Run() {
	locationIDsleft, locationIDright := readInput()

	star1, err := calculateTotalDistances(locationIDsleft, locationIDright)
	if err != nil {
		log.Fatalln("error calculating distance of IDs")
	}

	fmt.Println("Result day 1 star 1: ", star1)

	star2, err := calculateSimilarityScore(locationIDsleft, locationIDright)
	if err != nil {
		log.Fatalln("error calculating similarity score")
	}

	fmt.Println("Result day 1 star 2: ", star2)
}
