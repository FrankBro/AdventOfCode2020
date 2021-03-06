package day7_test

import (
	"testing"

	"github.com/FrankBro/AdventOfCode2020/day7"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	output, err := day7.Part1()
	require.NoError(t, err)
	require.EqualValues(t, 151, output)
}

func TestPart2(t *testing.T) {
	t.Parallel()
	output, err := day7.Part2()
	require.NoError(t, err)
	require.EqualValues(t, 41559, output)
}
