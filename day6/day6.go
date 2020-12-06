package day6

import (
	"fmt"

	"github.com/FrankBro/AdventOfCode2020/util"
)

const (
	errReadLines = "failed to read input as lines: %w"
)

func union(group []map[rune]struct{}) map[rune]struct{} {
	union := group[0]
	for _, person := range group[1:] {
		for answer := range person {
			union[answer] = struct{}{}
		}
	}
	return union
}

func intersection(group []map[rune]struct{}) map[rune]struct{} {
	intersection := group[0]
	for _, person := range group[1:] {
		for answer := range intersection {
			if _, ok := person[answer]; !ok {
				delete(intersection, answer)
			}
		}
	}
	return intersection
}

func group(lines []string) [][]map[rune]struct{} {
	groups := make([][]map[rune]struct{}, 0)
	group := make([]map[rune]struct{}, 0)
	for _, line := range lines {
		if line == "" {
			groups = append(groups, group)
			group = make([]map[rune]struct{}, 0)
		} else {
			answers := make(map[rune]struct{})
			for _, answer := range line {
				answers[answer] = struct{}{}
			}
			group = append(group, answers)
		}
	}
	return groups
}

func Part1() (int, error) {
	lines, err := util.ReadLines()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	groups := group(lines)
	var count int
	for _, group := range groups {
		union := union(group)
		count += len(union)
	}
	return count, nil
}

func Part2() (int, error) {
	lines, err := util.ReadLines()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	groups := group(lines)
	var count int
	for _, group := range groups {
		intersection := intersection(group)
		count += len(intersection)
	}
	return count, nil
}
