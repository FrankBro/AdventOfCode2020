package day4_test

import (
	"testing"

	"github.com/FrankBro/AdventOfCode2020/day4"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	output, err := day4.Part1()
	require.NoError(t, err)
	require.EqualValues(t, 247, output)
}

func TestPart2(t *testing.T) {
	t.Parallel()
	output, err := day4.Part2()
	require.NoError(t, err)
	require.EqualValues(t, 145, output)
}
