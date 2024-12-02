package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/funkymcb/AdventOfCode/internal/day1"
)

const (
	YEAR      = 2024
	INPUT_DIR = "inputs"
)

func getInputFile(day int, filePath string) error {
	url := fmt.Sprintf("http://adventofcode.com/%d/day/%d/input", YEAR, day)

	sessionCookie := os.Getenv("AOC_SESSION")
	if sessionCookie == "" {
		return fmt.Errorf("session cookie is empty. make sure to export AOC_SESSION")
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error performing get request against: %s", url)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return fmt.Errorf("\nhttp status: %d,\nmessage: %s\n", resp.StatusCode, string(body))
	}

	fmt.Println("Successfully fetched input content. Writing to file:", filePath)
	if err = os.WriteFile(filePath, body, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func checkInputFile(day int) error {
	inputFileName := fmt.Sprintf("day_%d.txt", day)
	inputFilePath := filepath.Join(INPUT_DIR, inputFileName)

	// check if file already exists
	if _, err := os.Stat(inputFilePath); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Input file %s does not exist yet. Fetching...\n", inputFileName)

		// get input file
		if err := getInputFile(day, inputFilePath); err != nil {
			return err
		}

		return nil
	}

	return nil
}

func main() {
	dayPtr := flag.Int("day", 0, "Valid inputs: 1 - 24. the day to be executed.")
	flag.Parse()

	day := *dayPtr

	if day <= 0 || day > 24 {
		fmt.Println("pls specify the -day flag (1 - 24). Use -h for help")
		os.Exit(1)
	}

	if err := checkInputFile(day); err != nil {
		fmt.Println("error checking for input file", err)
		os.Exit(1)
	}

	day1.Run()
}
