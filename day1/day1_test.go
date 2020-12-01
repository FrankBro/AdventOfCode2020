package day1_test

import (
	"testing"

	"github.com/FrankBro/AdventOfCode2020/day1"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	require.EqualValues(t, 712075, day1.Part1())
}

func TestPart2(t *testing.T) {
	t.Parallel()
	require.EqualValues(t, 145245270, day1.Part2())
}
