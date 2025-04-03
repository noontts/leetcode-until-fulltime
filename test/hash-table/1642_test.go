package hashtable_test

import (
	hash_table "leetcode-tkankath/solutions/hash-table"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	testCase := []struct {
		input  string
		expect int
	}{
		{"aa", 0},
		{"abca", 2},
		{"cbzxy", -1},
		{"cabbac", 4},
		{"mgntdygtxrvxjnwksqhxuxtrv", 18},
	}

	for i, tc := range testCase {
		t.Run("MaxLengthBetweenEqualCharacters case: "+strconv.Itoa(i+1), func(t *testing.T) {
			t.Parallel()
			result := hash_table.MaxLengthBetweenEqualCharacters(tc.input)
			assert.Equal(t, tc.expect, result)
		})
	}
}
