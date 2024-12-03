package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
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

	fmt.Printf("Part1: %d\n", part1(puzzle_input))
	fmt.Printf("Part2: %d\n", part2(puzzle_input))
}

func part1(puzzle_input []string) int {
	total_safe := 0

	// Do something with puzzle input
	for _, line := range puzzle_input {
		report_str := strings.Fields(line)
		report := arr_str_to_arr_int(report_str)
		if determine_safe(report) {
			total_safe++
		}
	}
	return total_safe
}

func part2(puzzle_input []string) int {
	total_safe := 0

	for _, line := range puzzle_input {
		report_str := strings.Fields(line)
		report := arr_str_to_arr_int(report_str)
		if determine_safe(report) {
			total_safe++
		} else if safety_dampener(report) {
			total_safe++
		}
	}
	return total_safe
}

func arr_str_to_arr_int(arr []string) []int {
	ret := make([]int, 0)
	for _, val := range arr {
		num, _ := strconv.Atoi(val)
		ret = append(ret, num)
	}
	return ret
}

// Gets the difference between each adjacent number in an array and appends it to another array. Returns that.
func get_arr_diffs(arr []int) []int {
	arr_diffs := make([]int, 0)
	for i := 0; i < len(arr)-1; i++ {
		arr_diffs = append(arr_diffs, int(math.Abs(float64(arr[i]-arr[i+1]))))
	}

	return arr_diffs
}

func determine_safe(report []int) bool {
	// We can check if each array is sorted. Either ascending or descending.
	//If it is, then we just need to check diffs.
	sortedAsc := sort.SliceIsSorted(report, func(p, q int) bool {
		return report[p] > report[q]
	})

	// fmt.Println(sortedAsc)

	sortedDesc := sort.SliceIsSorted(report, func(p, q int) bool {
		return report[p] < report[q]
	})

	// fmt.Println(sortedDesc)

	//this is kinda janky
	// appropriate_levels := true
	// if sortedAsc || sortedDesc {
	// 	diffs := get_arr_diffs(report)
	// 	// fmt.Println(diffs)
	// 	for i := 0; i < len(diffs); i++ {
	// 		if diffs[i] == 0 || int(math.Abs(float64(diffs[i]))) >= 3 {
	// 			appropriate_levels = false
	// 		}
	// 	}
	// } else {
	// 	appropriate_levels = false
	// }
	//
	// if appropriate_levels {
	// 	total_safe++
	// }

	variance := get_arr_diffs(report)
	appropriate_variance := true
	if slices.Max(variance) > 3 || slices.Min(variance) == 0 {
		appropriate_variance = false
	}

	if (sortedAsc || sortedDesc) && appropriate_variance {
		return true
	}
	return false
}

func safety_dampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		var new_arr []int
		if i == 0 {
			new_arr = report[1:]
		} else if i == len(report)-1 {
			new_arr = report[:len(report)-1]
		} else {
			new_arr = append(new_arr, report[:i]...)
			new_arr = append(new_arr, report[i+1:]...)
		}
		if determine_safe(new_arr) {
			return true
		}
	}
	return false
}
