package day1

import (
	"errors"
	"fmt"

	"github.com/FrankBro/AdventOfCode2020/util"
)

const (
	addedValue     = 2020
	errReadNumbers = "failed to read input as numbers: %w"
)

var ErrNotFound = errors.New("not found")

func Part1() (int, error) {
	numbers, err := util.ReadNumbers()
	if err != nil {
		return 0, fmt.Errorf(errReadNumbers, err)
	}
	for i, ni := range numbers {
		for j, nj := range numbers {
			if i == j {
				continue
			}
			if ni+nj == addedValue {
				return ni * nj, nil
			}
		}
	}
	return 0, ErrNotFound
}

func Part2() (int, error) {
	numbers, err := util.ReadNumbers()
	if err != nil {
		return 0, fmt.Errorf(errReadNumbers, err)
	}
	for i, ni := range numbers {
		for j, nj := range numbers {
			for k, nk := range numbers {
				if i == j || i == k || j == k {
					continue
				}
				if ni+nj+nk == addedValue {
					return ni * nj * nk, nil
				}
			}
		}
	}
	return 0, ErrNotFound
}
