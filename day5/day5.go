package day5

import (
	"fmt"

	"github.com/FrankBro/AdventOfCode2020/util"
)

const (
	errReadLines = "failed to read input as lines: %w"
	maxRow       = 127
	maxCol       = 7
	multRow      = 8
	two          = 2 // Sometimes I think the magic number linters is overboard ...
)

func findMiddle(min, max int, add, sub rune, chars string) int {
	for _, char := range chars {
		diff := (max - min + 1) / two
		switch char {
		case add:
			min += diff
		case sub:
			max -= diff
		}
	}
	return min
}

func seatID(row, col int) int {
	return row*multRow + col
}

func parse(line string) int {
	row := findMiddle(0, maxRow, 'B', 'F', line[:7])
	col := findMiddle(0, maxCol, 'R', 'L', line[7:])
	seat := seatID(row, col)
	return seat
}

func Part1() (int, error) {
	lines, err := util.ReadLinesNoEmpty()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	var highest int
	for _, line := range lines {
		seat := parse(line)
		if seat > highest {
			highest = seat
		}
	}
	return highest, nil
}

func Part2() (int, error) {
	lines, err := util.ReadLinesNoEmpty()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	seats := make(map[int]struct{})
	for _, line := range lines {
		seat := parse(line)
		seats[seat] = struct{}{}
	}
	for row := 1; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			seat := seatID(row, col)
			if _, ok := seats[seat]; !ok {
				_, before := seats[seat-1]
				_, after := seats[seat+1]
				if before && after {
					return seat, nil
				}
			}
		}
	}
	return 0, nil
}
