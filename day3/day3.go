package day3

import (
	"fmt"

	"github.com/FrankBro/AdventOfCode2020/util"
)

const (
	errReadLines = "failed to read input as lines: %w"
)

func getTrees(lines []string, dx, dy int) int {
	var x, y int
	var trees int
	size := len(lines[0])
	for {
		x += dx
		y += dy
		relx := x % size
		if lines[y][relx] == '#' {
			trees++
		}
		if y == len(lines)-1 {
			break
		}
	}
	return trees
}

func Part1() (int, error) {
	lines, err := util.ReadLinesNoEmpty()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	trees := getTrees(lines, 3, 1)
	return trees, nil
}

func Part2() (int, error) {
	lines, err := util.ReadLinesNoEmpty()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	t1 := getTrees(lines, 1, 1)
	t2 := getTrees(lines, 3, 1)
	t3 := getTrees(lines, 5, 1)
	t4 := getTrees(lines, 7, 1)
	t5 := getTrees(lines, 1, 2)
	return t1 * t2 * t3 * t4 * t5, nil
}
