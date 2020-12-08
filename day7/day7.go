package day7

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/FrankBro/AdventOfCode2020/util"
)

const (
	errReadLines = "failed to read input as lines: %w"
	containToken = " bags contain "
	commaToken   = ", "
	periodToken  = "."
)

var errInvalidLine = errors.New("invalid line")

type color = string

func parseLines(lines []string) (map[color]map[color]int, error) {
	bags := make(map[string]map[color]int, len(lines))
	for _, line := range lines {
		color, contains, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		bags[color] = contains
	}
	return bags, nil
}

func parseLine(line string) (string, map[color]int, error) {
	index := strings.Index(line, containToken)
	if index == -1 {
		return "", nil, errInvalidLine
	}
	color := line[:index]
	line = line[index+len(containToken):]
	count := strings.Count(line, ", ") + 1
	contains := make(map[string]int, count)
	for i := 0; i < count; i++ {
		var contained string
		if i == count-1 {
			index = strings.Index(line, periodToken)
			if index == -1 {
				return "", nil, errInvalidLine
			}
			contained = line[:index]
		} else {
			index = strings.Index(line, commaToken)
			if index == -1 {
				return "", nil, errInvalidLine
			}
			contained = line[:index]
			line = line[index+len(commaToken):]
		}
		containedColor, containedCount, err := parseContained(contained)
		if err != nil {
			return "", nil, errInvalidLine
		}
		if containedCount != 0 {
			contains[containedColor] = containedCount
		}
	}
	return color, contains, nil
}

func parseContained(contained string) (string, int, error) {
	spaceBagIndex := strings.Index(contained, " bag")
	contained = contained[:spaceBagIndex]
	spaceIndex := strings.Index(contained, " ")
	maybeCount := contained[:spaceIndex]
	var count int
	if maybeCount != "no" {
		convertedCount, err := strconv.Atoi(maybeCount)
		if err != nil {
			return "", 0, errInvalidLine
		}
		count = convertedCount
	}
	color := contained[spaceIndex+1:]
	return color, count, nil
}

func contained(bags map[color]map[color]int, colors map[color]struct{}) int {
	count := len(colors)
	for {
		for bagColor, bagContains := range bags {
			for containedColor := range bagContains {
				if _, ok := colors[containedColor]; ok {
					colors[bagColor] = struct{}{}
				}
			}
		}
		if len(colors) != count {
			count = len(colors)
		} else {
			break
		}
	}
	return count - 1
}

func contain(bags map[color]map[color]int, color color) int {
	if contains, ok := bags[color]; ok {
		var in int
		for bagColor, bagCount := range contains {
			i := contain(bags, bagColor)
			in += i*bagCount + bagCount
		}
		return in
	}
	return 0
}

func Part1() (int, error) {
	lines, err := util.ReadLinesNoEmpty()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	bags, err := parseLines(lines)
	if err != nil {
		return 0, errInvalidLine
	}
	toFind := map[string]struct{}{
		"shiny gold": {},
	}
	count := contained(bags, toFind)
	return count, nil
}

func Part2() (int, error) {
	lines, err := util.ReadLinesNoEmpty()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	bags, err := parseLines(lines)
	if err != nil {
		return 0, errInvalidLine
	}
	return contain(bags, "shiny gold"), nil
}
