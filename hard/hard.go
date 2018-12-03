package heard

import "runtime/debug"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 翻转单链表中每个 k 长度的节点
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	pNode := head
	temp := k
	for temp > 0 {
		pNode = pNode.Next
		temp--
		if pNode == nil && temp > 0 {
			return head
		}
	}
	temp = k
	curNode := head
	preNode := new(ListNode)
	nextNode := new(ListNode)
	for temp > 0 && curNode != nil {
		nextNode = curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = nextNode
		temp--
	}
	head.Next = reverseKGroup(nextNode, k)
	return preNode
}

/**
  1 -> 2 -> 3 -> 4
  1 <- 2 <- 3 <- 4
*/
func ReverseList(head *ListNode) *ListNode { // 翻转单链表
	pNode := head
	if pNode == nil {
		return nil
	}
	var preNode *ListNode // 前驱
	preNode = nil
	for pNode.Next != nil {
		nextNode := pNode.Next // 后继
		pNode.Next = preNode
		preNode = pNode
		pNode = nextNode
	}
	return pNode
}

func FindSubstring(s string, words []string) []int {
	return nil
}

/**
寻找下一个更大的序列
1,2,3 -> 1,3,2
3,2,1 -> 1,2,3
1,1,5 -> 1,5,1
*/
func NextPermutation(nums []int) {

	if nums == nil || len(nums) <= 1 {
		return
	}
	index := len(nums) - 2
	for ; index >= 0; index-- {
		if nums[index] < nums[index+1] {
			break
		}
	}
	if index < 0 {
		ReverseSlice(nums)
		return
	} else {
		for end := len(nums) - 1; end > index; end-- {
			if nums[index] < nums[end] {
				nums[index], nums[end] = nums[end], nums[index]
				ReverseSlice(nums[index+1:])
				break
			}

		}
		return
	}
}

func ReverseSlice(nums []int) {
	start := 0
	end := len(nums) - 1
	for end > start {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func LongestValidParentheses(s string) int {

}
