package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	// Do something with puzzle input
	fmt.Printf("Part1: %d\n", part1(puzzle_input))
	fmt.Printf("Part2: %d\n", part2(puzzle_input))
}

func part1(puzzle_input []string) int {
	total := 0
	for _, line := range puzzle_input {
		regex := regexp.MustCompile(`(?m)mul\((?P<firstnum>\d+),(?P<secondnum>\d+)\)`)
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			total += atoi(match[1]) * atoi(match[2])
		}
	}

	return total
}

func part2(puzzle_input []string) int {
	var final_split []string

	line := strings.Join(puzzle_input, " ")
	// We can split the string on do() and this will give us <enabled instructions>don't()<disabled instructions>
	// This then allows us to just split on don't() and only handle the first index of each.
	split1 := regexp.MustCompile(`(?m)do\(\)`)
	split2 := regexp.MustCompile(`(?m)don\'t\(\)`)

	do_split := split1.Split(line, -1)
	for i := 0; i < len(do_split); i++ {
		dont_split := split2.Split(do_split[i], 2)
		final_split = append(final_split, dont_split[0])
	}
	// I think we can just call part1 now on the new data.
	total := part1(final_split)
	return total
}

func atoi(num_str string) int {
	num, _ := strconv.Atoi(num_str)
	return num
}
