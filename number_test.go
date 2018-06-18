package agnumber

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestI64toa(t *testing.T) {
	require.Equal(t, "12", I64toa(12))
}

func TestI64tox(t *testing.T) {
	require.Equal(t, "c", I64tox(12))
	require.Equal(t, "1f", I64tox(31))
}

func TestF64toa(t *testing.T) {
	require.Equal(t, "12.56789", F64toa(12.56789))
	require.Equal(t, "12.57", F64toa(12.56789, 2))
	require.Equal(t, "12.5678900000", F64toa(12.56789, 10))
}

func TestAtoi64(t *testing.T) {
	n, err := Atoi64("12")
	require.Nil(t, err)
	require.Equal(t, int64(12), n)

	_, err = Atoi64("12.0")
	require.NotNil(t, err)

	_, err = Atoi64("12,0")
	require.NotNil(t, err)

	_, err = Atoi64("12c")
	require.NotNil(t, err)
}

func TestAtof64(t *testing.T) {
	n, err := Atof64("12")
	require.Nil(t, err)
	require.Equal(t, 12.0, n)

	n, err = Atof64("12.0")
	require.Nil(t, err)
	require.Equal(t, 12.0, n)

	_, err = Atof64("12,0")
	require.NotNil(t, err)

	_, err = Atof64("12c")
	require.NotNil(t, err)
}

func TestMinMax(t *testing.T) {
	testCases := []struct {
		input    []int
		min, max int
	}{
		{nil, math.MinInt64, math.MaxInt64},
		{[]int{}, math.MinInt64, math.MaxInt64},
		{[]int{5, 8, 45, 1, 78}, 1, 78},
		{[]int{1, 5, 8, 1, 1}, 1, 8},
		{[]int{1, 5, 8, -1, 1}, -1, 8},
		{[]int{1, 8, 8, -1, -1}, -1, 8},
	}

	for _, testCase := range testCases {
		require.Equalf(t, testCase.min, Min(testCase.input...), "input %v", testCase.input)
		require.Equalf(t, testCase.max, Max(testCase.input...), "input %v", testCase.input)
	}
}
