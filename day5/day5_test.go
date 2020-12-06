package day5_test

import (
	"testing"

	"github.com/FrankBro/AdventOfCode2020/day5"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	output, err := day5.Part1()
	require.NoError(t, err)
	require.EqualValues(t, 896, output)
}

func TestPart2(t *testing.T) {
	t.Parallel()
	output, err := day5.Part2()
	require.NoError(t, err)
	require.EqualValues(t, 659, output)
}
