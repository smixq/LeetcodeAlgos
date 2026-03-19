package main

// Definition for singly-linked list.
type ListNode struct {
  Val int
Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	var currentList *ListNode = &ListNode{0, nil}
	var hadList *ListNode = currentList
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			currentList.Next = list1
			list1 = list1.Next
		} else {
			currentList.Next = list2
			list2 = list2.Next
		}
		currentList = currentList.Next
	}
	if list1 == nil {
		currentList.Next = list2
	} else if list2 == nil {
		currentList.Next = list1
	} else {
		currentList.Next = nil
	}

	return hadList.Next
}
