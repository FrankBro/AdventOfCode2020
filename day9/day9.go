package day9

import (
	"fmt"
	"math"

	"github.com/FrankBro/AdventOfCode2020/util"
)

const (
	errReadLines = "failed to read input as lines: %w"
	preamble     = 25
)

func noSumNumber(count int, numbers []int) int {
	for i := count; i < len(numbers); i++ {
		var found bool
		number := numbers[i]
		for j := i - count; j < i; j++ {
			for k := i - count; k < i; k++ {
				if j == k {
					continue
				}
				if numbers[j]+numbers[k] == number {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return number
		}
	}
	return 0
}

func Part1() (int, error) {
	numbers, err := util.ReadNumbers()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	final := noSumNumber(preamble, numbers)
	return final, nil
}

func Part2() (int, error) {
	numbers, err := util.ReadNumbers()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	special := noSumNumber(preamble, numbers)
	var done bool
	var minIndex, maxIndex int
	for i, number := range numbers {
		acc := number
		for j := i + 1; j < len(numbers); j++ {
			acc += numbers[j]
			if acc == special {
				minIndex = i
				maxIndex = j
				done = true
				break
			}
			if acc > special {
				break
			}
		}
		if done {
			break
		}
	}
	min := math.MaxInt64
	max := 0
	for i := minIndex; i <= maxIndex; i++ {
		number := numbers[i]
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}
	return min + max, nil
}
