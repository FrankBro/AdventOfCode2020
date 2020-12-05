package day5

import (
	"fmt"

	"github.com/FrankBro/AdventOfCode2020/util"
)

const (
	errReadLines = "failed to read input as lines: %w"
)

func Part1() (int, error) {
	lines, err := util.ReadLines()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	return len(lines), nil
}

func Part2() (int, error) {
	lines, err := util.ReadLines()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	return len(lines), nil
}
