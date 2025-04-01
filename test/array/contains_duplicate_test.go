package array_test

import (
	"leetcode-tkankath/solutions/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testCase := []struct {
		input  []int
		expect bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tc := range testCase {
		t.Parallel()

		result := array.ContainsDuplicate(tc.input)
		assert.Equal(t, tc.expect, result, "they should be equal")
	}
}
