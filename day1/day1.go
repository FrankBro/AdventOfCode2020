package day1

import (
	"github.com/FrankBro/AdventOfCode2020/util"
)

const (
	addedValue = 2020
)

func Part1() int {
	numbers := util.ReadNumbers()
	for i, ni := range numbers {
		for j, nj := range numbers {
			if i == j {
				continue
			}
			if ni+nj == addedValue {
				return ni * nj
			}
		}
	}
	panic("not found")
}

func Part2() int {
	numbers := util.ReadNumbers()
	for i, ni := range numbers {
		for j, nj := range numbers {
			for k, nk := range numbers {
				if i == j || i == k || j == k {
					continue
				}
				if ni+nj+nk == addedValue {
					return ni * nj * nk
				}
			}
		}
	}
	panic("not found")
}
