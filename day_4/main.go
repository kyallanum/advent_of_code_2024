package main

import (
	"bufio"
	"fmt"
	"os"
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

	fmt.Printf("Part1: %d\n", part1(puzzle_input))
	fmt.Printf("Part2: %d\n", part2(puzzle_input))

}

func part1(puzzle_input []string) int {

	// We need to transform puzzle_input into [][]string
	xmas_message := make([][]string, 0)
	for _, line := range puzzle_input {
		xmas_message = append(xmas_message, strings.Split(line, ""))
	}

	//Each direction forward, backward, diagonal
	directions := [][]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}

	total_xmas := 0

	for row_idx := 0; row_idx < len(xmas_message); row_idx++ {
		for col_idx := 0; col_idx < len(xmas_message[row_idx]); col_idx++ {
			// We first find a position that has X
			if xmas_message[row_idx][col_idx] != "X" {
				continue
			}

			// We then find an M. If we do, we have that direction and we stick with it.
			for _, direction := range directions {
				// We are just checking bounds here.
				if (row_idx+direction[0] < 0) || (row_idx+direction[0] >= len(xmas_message)) {
					continue
				}
				if (col_idx+direction[1] < 0) || (col_idx+direction[1] >= len(xmas_message[row_idx])) {
					continue
				}
				//Check if that value is "M"
				if xmas_message[row_idx+direction[0]][col_idx+direction[1]] != "M" {
					continue
				}
				// We just continue in the same direction by multiplying it by 2
				if (row_idx+(direction[0]*2) < 0) || (row_idx+(direction[0]*2) >= len(xmas_message)) {
					continue
				}
				if (col_idx+(direction[1]*2) < 0) || (col_idx+(direction[1]*2) >= len(xmas_message[row_idx])) {
					continue
				}
				if xmas_message[row_idx+(direction[0]*2)][col_idx+(direction[1]*2)] != "A" {
					continue
				}
				// Same  here
				if (row_idx+(direction[0]*3) < 0) || (row_idx+(direction[0]*3) >= len(xmas_message)) {
					continue
				}
				if (col_idx+(direction[1]*3) < 0) || (col_idx+(direction[1]*3) >= len(xmas_message[row_idx])) {
					continue
				}
				if xmas_message[row_idx+(direction[0]*3)][col_idx+(direction[1]*3)] != "S" {
					continue
				}
				// If all tests have passed then we increment and continue.
				total_xmas++
			}
		}
	}
	return total_xmas
}

func part2(puzzle_input []string) int {
	xmas_message := make([][]string, 0)
	for _, line := range puzzle_input {
		xmas_message = append(xmas_message, strings.Split(line, ""))
	}

	xmas_total := 0
	// The A's are the only unique thing in the X. so we search for that.
	directions := [][][]int{
		{
			{1, 1}, {1, -1}, // M's are both down
		},
		{
			{1, 1}, {-1, 1}, // M's are to the right
		},
		{
			{-1, -1}, {-1, 1}, // M's are both up
		},
		{
			{-1, -1}, {1, -1}, // M's are both to the left
		},
	}

	// We only have to search the inner ring because the X always has the same shape
	for row_idx := 1; row_idx < len(xmas_message)-1; row_idx++ {
		for col_idx := 1; col_idx < len(xmas_message[row_idx])-1; col_idx++ {
			if xmas_message[row_idx][col_idx] != "A" {
				continue
			}

			for _, direction := range directions {
				if xmas_message[row_idx+direction[0][0]][col_idx+direction[0][1]] != "M" {
					continue
				}
				if xmas_message[row_idx+direction[1][0]][col_idx+direction[1][1]] != "M" {
					continue
				}
				if xmas_message[row_idx+(direction[0][0]*-1)][col_idx+(direction[0][1]*-1)] != "S" {
					continue
				}
				if xmas_message[row_idx+(direction[1][0]*-1)][col_idx+(direction[1][1]*-1)] != "S" {
					continue
				}
				xmas_total++
				break
			}
		}
	}

	return xmas_total
}
