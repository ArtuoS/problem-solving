package mergetwosortedlists

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	getNumbers := func(node *ListNode) []int {
		if node == nil {
			return []int{}
		}

		arr := []int{node.Val}
		current := node
		for {
			if current.Next == nil {
				break
			}

			current = current.Next
			arr = append(arr, current.Val)
		}
		return arr
	}

	mergedArr := []int{}
	mergedArr = append(mergedArr, getNumbers(list1)...)
	mergedArr = append(mergedArr, getNumbers(list2)...)
	sort.Ints(mergedArr)

	var node *ListNode
	for _, v := range mergedArr {
		if node == nil {
			node = &ListNode{
				Val: v,
			}
			continue
		}

		curr := node
		for ; curr.Next != nil; curr = curr.Next {
		}
		curr.Next = &ListNode{
			Val: v,
		}
	}

	return node
}
