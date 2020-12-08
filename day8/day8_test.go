package day8_test

import (
	"testing"

	"github.com/FrankBro/AdventOfCode2020/day8"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	output, err := day8.Part1()
	require.NoError(t, err)
	require.EqualValues(t, 1420, output)
}

func TestPart2(t *testing.T) {
	t.Parallel()
	output, err := day8.Part2()
	require.NoError(t, err)
	require.EqualValues(t, 1245, output)
}
