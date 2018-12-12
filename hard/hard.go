package heard

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
