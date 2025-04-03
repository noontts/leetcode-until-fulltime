package main_test

import (
	"leetcode-tkankath/solutions"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelfDividingNumbers(t *testing.T) {
	testCase := []struct {
		left   int
		right  int
		expect []int
	}{
		{1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		{47, 85, []int{48, 55, 66, 77}},
	}

	for i, tc := range testCase {
		t.Run("Self dividing numbers case: "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			result := solutions.SelfDividingNumbers(tc.left, tc.right)
			assert.Equal(t, tc.expect, result, "they should be equal")
		})
	}
}
