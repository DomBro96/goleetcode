package heard

import (
	"sort"
)

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

type Stack []int

func (stack Stack) Len() int {

	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack *Stack) Push(value int) {
	*stack = append(*stack, value)
}

func (stack Stack) Top() int {
	if len(stack) == 0 {
		return -1
	}
	return stack[len(stack)-1]
}

func (stack *Stack) Pop() int {
	theStack := *stack
	if len(theStack) == 0 {
		return -1
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value
}

func (stack Stack) isEmpty() bool {
	return len(stack) == 0
}

//32. Longest Valid Parentheses
func LongestValidParentheses(s string) int {
	n := len(s)
	longest := 0
	var stack Stack
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack.Push(i)
		} else {
			if !stack.isEmpty() {
				if s[stack.Top()] == '(' {
					stack.Pop()
				} else {
					stack.Push(i)
				}
			} else {
				stack.Push(i)
			}
		}
	}
	if stack.isEmpty() {
		longest = n
	} else {
		a := n
		b := 0
		for !stack.isEmpty() {
			b = stack.Pop()
			longest = max(longest, a-b-1)
			a = b
		}
		longest = max(longest, a)
	}
	return longest
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

// 37
func SolveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				var filler byte
				for filler = '1'; filler <= '9'; filler++ {
					if isValid(board, i, j, filler) {
						board[i][j] = filler
						if solve(board) {
							return true
						}
					}
				}
				board[i][j] = '.' // backtracking
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, filler byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == filler {
			return false
		}
		if board[i][col] == filler {
			return false
		}
		rowIndex := 3 * (row / 3)
		colIndex := 3 * (col / 3)
		if board[rowIndex+i/3][colIndex+i%3] == filler {
			return false
		}
	}
	return true
}

// 41
func FirstMissingPositive(nums []int) int {
	if len(nums) < 1 || nums == nil {
		return 1
	}
	sort.Ints(nums)
	min := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] == min {
			min++
		}
	}
	return min
}

// 42
func Trap(height []int) int {
	if len(height) == 0 || height == nil {
		return 0
	}
	leftIndex, rightIndex := 0, len(height)-1
	leftMax, rightMax, result := height[leftIndex], height[rightIndex], 0
	for leftIndex < rightIndex {
		if leftMax <= rightMax {
			leftIndex++
			if height[leftIndex] < leftMax {
				result += leftMax - height[leftIndex]
			} else {
				leftMax = height[leftIndex]
			}
		} else {
			rightIndex--
			if height[rightIndex] < rightMax {
				result += rightMax - height[rightIndex]
			} else {
				rightMax = height[rightIndex]
			}
		}
	}
	return result
}

// 44 dynamic programming
func IsMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) > 0 && len(p) == 0 {
		return false
	}
	var dp [][]bool
	for i := 0; i <= len(s); i++ {
		dp = append(dp, make([]bool, len(p)+1))
	}
	// initialization memo
	dp[0][0] = true
	for i := 1; i <= len(p); i++ {
		if p[i-1] != '*' {
			dp[0][i] = false
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// dp[i - 1][j]: '*' match n time; dp[i][j - 1]: '*' match 0 time;
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[len(s)][len(p)]
}
