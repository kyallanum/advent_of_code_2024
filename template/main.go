package main

import (
	"bufio"
	"fmt"
	"os"
)

func load_file(file_path string) ([]string, error) {
	file, err := os.Open(file_path)

	if err != nil {
		return nil, fmt.Errorf("Could not open file: %s - %v", file, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	file_contents := make([]string, 0)

	for scanner.Scan() {
		file_contents = append(file_contents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("An error occurred with the scanner: %v", err)
	}
	return file_contents, nil
}

func main() {
	puzzle_input, err := load_file("input.txt")
	if err != nil {
		panic(err)
	}

	// Do something with puzzle input
	for line := range puzzle_input {
		fmt.Println(line)
	}
}
