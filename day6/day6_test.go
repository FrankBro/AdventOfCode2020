package day6_test

import (
	"testing"

	"github.com/FrankBro/AdventOfCode2020/day6"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	output, err := day6.Part1()
	require.NoError(t, err)
	require.EqualValues(t, 6612, output)
}

func TestPart2(t *testing.T) {
	t.Parallel()
	output, err := day6.Part2()
	require.NoError(t, err)
	require.EqualValues(t, 3268, output)
}
