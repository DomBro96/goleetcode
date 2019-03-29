package heard

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{3, nil}
	node4 := ListNode{4, nil}
	node5 := ListNode{5, nil}
	node6 := new(ListNode)
	node6.Val = 6
	node6.Next = nil
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = node6
	newHead := ReverseList(&node1)
	fmt.Println(newHead.Val)
}

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	nums1 := []int{3, 2, 1}
	nums2 := []int{1, 3, 5, 4, 3}
	NextPermutation(nums)
	NextPermutation(nums1)
	NextPermutation(nums2)
	fmt.Println(nums)
	fmt.Println(nums1)
	fmt.Println(nums2)
}

func TestLongestValidParentheses(t *testing.T) {
	s := "(()"
	fmt.Println(LongestValidParentheses(s))
	s1 := ")()())"
	fmt.Println(LongestValidParentheses(s1))
	s2 := ")(()()))"
	fmt.Println(LongestValidParentheses(s2))
	s3 := "())"
	fmt.Println(LongestValidParentheses(s3))
}

func TestSolveSudoku(t *testing.T) {
	input := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(input)
	SolveSudoku(input)
	fmt.Println(input)
}

func TestFirstMissingPositive(t *testing.T) {
	input := []int{1, 2, 0}
	input1 := []int{3, 4, -1, 1}
	input2 := []int{7, 8, 9, 11, 12}
	input3 := []int{2, 1}
	fmt.Println(FirstMissingPositive(input))
	fmt.Println(FirstMissingPositive(input1))
	fmt.Println(FirstMissingPositive(input2))
	fmt.Println(FirstMissingPositive(input3))

}

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(Trap(height))
	height1 := []int{2, 0, 2}
	fmt.Println(Trap(height1))
	height2 := []int{4, 2, 3}
	fmt.Println(Trap(height2))
}

func TestIsMatch(t *testing.T) {
	fmt.Println(IsMatch("aa", "a"))
	fmt.Println(IsMatch("aa", "*"))
	fmt.Println(IsMatch("cb", "?a"))
	fmt.Println(IsMatch("adceb", "*a*b"))
	fmt.Println(IsMatch("acdcb", "a*c?b"))
	fmt.Println(IsMatch("ab", "?*"))
}

func TestMinDistance(t *testing.T) {
	fmt.Println(MinDistance("horse", "ros"))
	fmt.Println(MinDistance("intention", "execution"))
	fmt.Println(MinDistance("sea", "eat"))
}
