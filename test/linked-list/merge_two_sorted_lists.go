package linkedlist_test

import (
	linkedlist "leetcode-tkankath/solutions/linked-list"
	"leetcode-tkankath/utils"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testCase := []struct {
		list1  *utils.ListNode
		list2  *utils.ListNode
		expect *utils.ListNode
	}{}

	for i, tc := range testCase {
		t.Run("Merge Two Sorted Lists case: "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			result := linkedlist.MergeTwoLists(tc.list1, tc.list2)
			assert.Equal(t, tc.expect, result, "they should be equal")
		})
	}
}
