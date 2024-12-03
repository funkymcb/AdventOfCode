package io

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/funkymcb/AdventOfCode/internal/config"
)

func ReadFile(day int) ([]string, error) {
	inputDir := config.INPUT_DIR
	inputFile := fmt.Sprintf("day_%d.txt", day)

	file, err := os.Open(filepath.Join(inputDir, inputFile))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}
