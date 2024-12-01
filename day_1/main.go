package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

	fmt.Println("Part 1:", part1(puzzle_input))
	fmt.Println("Part 2:", part2(puzzle_input))

}

func part1(input []string) int {
	// We put loc1 and loc2 in separate arrays, sort then get distance
	loc1 := make([]int, 0)
	loc2 := make([]int, 0)

	// Do something with puzzle input
	for _, line := range input {
		// Parse input
		locations := strings.Fields(line)
		loc1 = append(loc1, atoi(locations[0]))
		loc2 = append(loc2, atoi(locations[1]))
	}

	slices.Sort(loc1)
	slices.Sort(loc2)

	distance := 0
	for i := 0; i < len(loc1); i++ {
		distance += abs(loc1[i] - loc2[i])
	}

	return distance
}

func part2(input []string) int {
	// We make a map that counts num occurrences. Then reference each, multiply and add.
	loc1 := make([]int, 0)
	loc2 := make(map[int]int, 0)

	// Do something with puzzle input.
	for _, line := range input {
		locations := strings.Fields(line)
		loc1 = append(loc1, atoi(locations[0]))
		if _, success := loc2[atoi(locations[1])]; !success {
			loc2[atoi(locations[1])] = 1
		} else {
			loc2[atoi(locations[1])]++
		}
	}

	distance := 0
	for i := 0; i < len(loc1); i++ {
		cur_loc2 := 0
		cur_loc2 = loc2[loc1[i]]
		distance += loc1[i] * cur_loc2
	}
	return distance
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
