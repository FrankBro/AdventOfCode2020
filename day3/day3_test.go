package day3_test

import (
	"testing"

	"github.com/FrankBro/AdventOfCode2020/day3"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	output, err := day3.Part1()
	require.NoError(t, err)
	require.EqualValues(t, 176, output)
}

func TestPart2(t *testing.T) {
	t.Parallel()
	output, err := day3.Part2()
	require.NoError(t, err)
	require.EqualValues(t, 5872458240, output)
}
