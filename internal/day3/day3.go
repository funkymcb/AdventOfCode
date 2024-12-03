package day3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/funkymcb/AdventOfCode/internal/io"
)

func handleInput(day int) (string, error) {
	var memoryBuilder strings.Builder

	lines, err := io.ReadFile(day)
	if err != nil {
		return "", err
	}

	for _, line := range lines {
		_, err := memoryBuilder.WriteString(line)
		if err != nil {
			return "", err
		}
	}

	return memoryBuilder.String(), nil
}

func handleCorruptedMemory(memory string) (int, error) {
	match, err := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
	if err != nil {
		return 0, err
	}

	mulExpressions := match.FindAllString(memory, -1)

	var products []int
	for _, mul := range mulExpressions {
		m, err := regexp.Compile("[0-9]+")
		if err != nil {
			return 0, err
		}

		operants := m.FindAllString(mul, -1)
		o1, err := strconv.Atoi(operants[0])
		if err != nil {
			return 0, err
		}
		o2, err := strconv.Atoi(operants[1])
		if err != nil {
			return 0, err
		}

		products = append(products, (o1 * o2))
	}

	var result int
	for _, p := range products {
		result += p
	}

	return result, nil
}

// strippedMemory will remove everything after "don't()" up until the next "do()"
func stripMemory(memory string) (string, error) {
	matchDont, err := regexp.Compile("don't\\(\\)")
	if err != nil {
		return "", err
	}

	dontIndices := matchDont.FindAllStringIndex(memory, -1)

	matchDo, err := regexp.Compile("do\\(\\)")
	if err != nil {
		return "", err
	}

	doIndices := matchDo.FindAllStringIndex(memory, -1)

	var doBuffer int
	var strips []string
	for _, dont := range dontIndices {
		if dont[0] < doBuffer {
			continue
		}

		for _, do := range doIndices {
			if do[0] > dont[1] {
				//remove everything inbetween dont[0] and do[1]
				strips = append(strips, memory[dont[0]:do[1]])
				doBuffer = do[1]
				break
			} else {
				continue
			}
		}
	}

	// remove strips from memory
	// TODO not a good solution since there could theoretically be the equal strips
	for _, strip := range strips {
		memory = strings.Replace(memory, strip, "", -1)
	}

	return memory, nil
}

func handleExtendedMemory(memory string) (int, error) {
	strippedMemory, err := stripMemory(memory)
	if err != nil {
		return 0, err
	}

	result, err := handleCorruptedMemory(strippedMemory)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func Run(day int) {
	memory, err := handleInput(day)
	if err != nil {
		fmt.Println("error reading input", err)
	}

	star1, err := handleCorruptedMemory(memory)
	if err != nil {
		fmt.Println("error handling corrupted memory", err)
	}

	star2, err := handleExtendedMemory(memory)
	if err != nil {
		fmt.Println("error by extended memory handling", err)
	}

	io.PrintResult(day, star1, star2)
}
