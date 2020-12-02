package day2_test

import (
	"testing"

	"github.com/FrankBro/AdventOfCode2020/day2"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	output, err := day2.Part1()
	require.NoError(t, err)
	require.EqualValues(t, 536, output)
}

func TestPart2(t *testing.T) {
	t.Parallel()
	output, err := day2.Part2()
	require.NoError(t, err)
	require.EqualValues(t, 558, output)
}
