package linkedlist

import (
	"fmt"
	"leetcode-tkankath/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func MergeTwoLists(list1 *utils.ListNode, list2 *utils.ListNode) *utils.ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	resultListNode := &utils.ListNode{}
	currentListNode := resultListNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			currentListNode.Next = list1
			list1 = list1.Next
		} else {
			currentListNode.Next = list2
			list2 = list2.Next
		}

		currentListNode = currentListNode.Next
	}

	//If any have any one left node
	if list1 != nil {
		currentListNode.Next = list1
	} else {
		currentListNode.Next = list2
	}

	return resultListNode.Next
}

func logNode(node *utils.ListNode) {
	for node != nil {
		fmt.Printf("%v -> ", node.Val)

		node = node.Next
	}

	fmt.Println()
}
