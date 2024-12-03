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

func Run(day int) {
	memory, err := handleInput(day)
	if err != nil {
		fmt.Println("error reading input", err)
	}

	star1, err := handleCorruptedMemory(memory)
	if err != nil {
		fmt.Println("error handling corrupted memory", err)
	}

	io.PrintResult(day, star1, 0)
}
