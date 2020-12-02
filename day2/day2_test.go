package day2_test

import (
	"testing"

	"github.com/FrankBro/AdventOfCode2020/day2"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	require.EqualValues(t, 712075, day2.Part1())
}

func TestPart2(t *testing.T) {
	t.Parallel()
	require.EqualValues(t, 145245270, day2.Part2())
}
