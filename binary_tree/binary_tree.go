package binarytree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DepthFirstValues(root *TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	array := []*TreeNode{root}
	var currentNode *TreeNode
	for {
		if len(array) <= 0 {
			break
		}

		currentNode = array[len(array)-1]
		array = array[:len(array)-1]

		if currentNode != nil {
			result = append(result, currentNode.Val)
		}

		if currentNode.Right != nil {
			array = append(array, currentNode.Right)
		}

		if currentNode.Left != nil {
			array = append(array, currentNode.Left)
		}
	}

	return result
}

func DepthFirstValue2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{root.Val}
	result = append(result, DepthFirstValue2(root.Left)...)
	result = append(result, DepthFirstValue2(root.Right)...)

	fmt.Println(result)

	return result
}

func BreadthFirstValue(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	result := []int{}
	for {
		if len(queue) <= 0 {
			break
		}

		currentNode := queue[0]
		queue = queue[1:]
		result = append(result, currentNode.Val)
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}

		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}

	fmt.Println(result)
	return result
}

func TargetExists(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}
	for {
		if len(queue) <= 0 {
			break
		}

		currentNode := queue[0]
		queue = queue[1:]
		if currentNode.Val == target {
			return true
		}

		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}

		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}

	return false
}

func TreeSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	queue := []*TreeNode{root}
	for {
		if len(queue) <= 0 {
			break
		}

		currentNode := queue[0]
		queue = queue[1:]
		sum += currentNode.Val

		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}

		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}

	return sum
}

func TreeSum2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := root.Val
	sum += TreeSum2(root.Left)
	sum += TreeSum2(root.Right)

	return sum
}
