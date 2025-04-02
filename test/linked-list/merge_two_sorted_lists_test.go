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
	}{
		{utils.NewList([]int{1, 2, 4}), utils.NewList([]int{1, 3, 4}), utils.NewList([]int{1, 1, 2, 3, 4, 4})},
		{utils.NewList([]int{}), utils.NewList([]int{}), utils.NewList([]int{})},
		{utils.NewList([]int{}), utils.NewList([]int{0}), utils.NewList([]int{0})},
	}

	for i, tc := range testCase {
		t.Run("Merge Two Sorted Lists case: "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			result := linkedlist.MergeTwoLists(tc.list1, tc.list2)
			assert.True(t, utils.CompareLists(result, tc.expect), "List value must equal")
		})
	}
}
