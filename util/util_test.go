package util_test

import (
	"testing"

	"github.com/FrankBro/AdventOfCode2020/util"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Parallel()
	input := "1-3 a: abcde"
	parsed, err := util.Parse(input, "-", " ", ": ")
	require.NoError(t, err)
	expected := []string{"1", "3", "a", "abcde"}
	require.EqualValues(t, expected, parsed)
}
