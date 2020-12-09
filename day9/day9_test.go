package day9_test

import (
	"testing"

	"github.com/FrankBro/AdventOfCode2020/day9"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	output, err := day9.Part1()
	require.NoError(t, err)
	require.EqualValues(t, 88311122, output)
}

func TestPart2(t *testing.T) {
	t.Parallel()
	output, err := day9.Part2()
	require.NoError(t, err)
	require.EqualValues(t, 13549369, output)
}
