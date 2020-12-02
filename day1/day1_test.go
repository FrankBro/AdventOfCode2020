package day1_test

import (
	"testing"

	"github.com/FrankBro/AdventOfCode2020/day1"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	output, err := day1.Part1()
	require.NoError(t, err)
	require.EqualValues(t, 712075, output)
}

func TestPart2(t *testing.T) {
	t.Parallel()
	output, err := day1.Part2()
	require.NoError(t, err)
	require.EqualValues(t, 145245270, output)
}
