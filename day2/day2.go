package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FrankBro/AdventOfCode2020/util"
)

const (
	errReadLines = "failed to read input as lines: %w"
	errParse     = "failed to parse line: %w"
	errAtoi      = "failed to convert to integer: %w"
)

func getValues(line string) (int, int, string, string, error) {
	parsed, err := util.Parse(line, "-", " ", ": ")
	if err != nil {
		return 0, 0, "", "", fmt.Errorf(errParse, err)
	}
	first, err := strconv.Atoi(parsed[0])
	if err != nil {
		return 0, 0, "", "", fmt.Errorf(errAtoi, err)
	}
	second, err := strconv.Atoi(parsed[1])
	if err != nil {
		return 0, 0, "", "", fmt.Errorf(errAtoi, err)
	}
	letter := parsed[2]
	input := parsed[3]
	return first, second, letter, input, nil
}

func Part1() (int, error) {
	lines, err := util.ReadLinesNoEmpty()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	var valid int
	for _, line := range lines {
		min, max, letter, input, err := getValues(line)
		if err != nil {
			return 0, err
		}
		count := strings.Count(input, letter)
		isValid := count >= min && count <= max
		if isValid {
			valid++
		}
	}
	return valid, nil
}

func Part2() (int, error) {
	lines, err := util.ReadLinesNoEmpty()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	var valid int
	for _, line := range lines {
		first, second, letter, input, err := getValues(line)
		if err != nil {
			return 0, err
		}
		firstMatches := string(input[first-1]) == letter
		secondMatches := string(input[second-1]) == letter
		isValid := (firstMatches || secondMatches) && firstMatches != secondMatches
		if isValid {
			valid++
		}
	}
	return valid, nil
}
