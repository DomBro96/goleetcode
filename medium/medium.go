package medium

import (
	"fmt"
)

// 33. Search in Rotated Sorted Array
func Search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	i1 := 0
	i2 := n - 1 // 指向最小的元素
	im := 0
	// 找到第一个小于前一个数且小于后一个数的索引
	/**
	0 1 2 3 4 5
	3 4 5 0 1 2
	1 2 3 4 5 0
	*/
	for nums[i1] > nums[i2] {
		if i2-i1 == 1 {
			im = i2
			break
		}
		im = (i1 + i2) / 2
		if nums[im] >= nums[i1] {
			i1 = im
		} else if nums[im] <= nums[i2] {
			i2 = im
		}
	}

	if im == 0 {
		return BinSearch(nums, target)
	} else {
		if nums[im] == target {
			return im
		} else if nums[im] > target {
			return -1
		}
		nums1 := nums[:im]
		nums2 := nums[im:]
		if r := BinSearch(nums1, target); r != -1 {
			return r
		}
		if r := BinSearch(nums2, target); r != -1 {
			return im + r
		}
	}

	return -1
}

func BinSearch(nums []int, target int) int {
	var mid int
	end := len(nums) - 1
	start := 0
	for end >= start {
		mid = (start + end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 34 Find First and Last Position of Element in Sorted Array
func SearchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	begin := 0
	end := len(nums) - 1
	first := -1
	last := -1
	for end >= begin {
		mid := (begin + end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			begin = mid + 1
		} else {
			first = mid
			last = mid
			for first >= 0 && nums[first] == target {
				first--
			}
			first++
			for last <= len(nums)-1 && nums[last] == target {
				last++
			}
			last--
			break
		}
	}
	return []int{first, last}
}

func IsValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		rowsMap := make(map[byte]bool, 9)
		colsMap := make(map[byte]bool, 9)
		cubeMap := make(map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, exists := rowsMap[board[i][j]]; exists {
					return false
				} else {
					rowsMap[board[i][j]] = true
				}
			}
			if board[j][i] != '.' {
				if _, exists := colsMap[board[j][i]]; exists {
					return false
				} else {
					colsMap[board[j][i]] = true
				}
			}
			rowIndex := 3 * (i / 3)
			colIndex := 3 * (i % 3)
			if board[rowIndex+j/3][colIndex+j%3] != '.' {
				if _, exists := cubeMap[board[rowIndex+j/3][colIndex+j%3]]; exists {
					return false
				} else {
					cubeMap[board[rowIndex+j/3][colIndex+j%3]] = true
				}
			}
		}
	}
	return true
}

// 287
func FindDuplicate(nums []int) int {
	start := 1
	end := len(nums) - 1
	result := 0
	for {
		if start == end {
			result = start
			break
		}
		middle := (start + end) / 2
		if countNum(nums, start, middle) > middle-start+1 {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return result
}

func countNum(nums []int, start, end int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= start && nums[i] <= end {
			count++
		}
	}
	return count
}

//62 Unique Paths
func UniquePaths(m int, n int) int {
	//result := 0
	//doUniquePaths(n, m, 0, 0, m*n-1, &result)
	//return result
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([]int, m*n+1)
	dp[1] = 1
	for i := 2; i <= m; i++ {
		dp[i] = dp[i-1] + 0
	}
	return doUniquePath(m, dp)
}

// f(j) = sum{f(j - 1)+f(j-cols)}
func doUniquePath(cols int, dp []int) int {
	for i := cols + 1; i < len(dp); i++ {
		// first col don't have left path
		if i%cols == 1 || cols == 1 {
			dp[i] = dp[i-cols]
		} else {
			dp[i] = dp[i-1] + dp[i-cols]
		}
	}
	return dp[len(dp)-1]
}

// Recursive
func doUniquePaths(rows, cols, row, col, finish int, res *int) {
	if row == rows || col == cols {
		return
	}
	if (row*cols + col) == finish {
		*res++
		fmt.Printf("row: %d, col: %d, finish: %d\n", row, col, finish)
		return
	}
	doUniquePaths(rows, cols, row+1, col, finish, res)
	doUniquePaths(rows, cols, row, col+1, finish, res)
}

// 63 Unique Paths
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])
	if rows == 0 || cols == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	dp := make([][]int, 0)
	for i := 0; i <= rows; i++ {
		dp = append(dp, make([]int, cols+1))
	}
	dp[1][1] = 1
	// init
	for i := 2; i <= cols; i++ {
		if obstacleGrid[0][i-1] != 1 {
			dp[1][i] = dp[1][i-1]
		} else {
			dp[1][i] = 0
		}
	}
	return doUniquePathsWithObstacles(obstacleGrid, dp, rows, cols)
}

func doUniquePathsWithObstacles(obstacleGrid, dp [][]int, rows, cols int) int {
	for i := 2; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if obstacleGrid[i-1][j-1] == 1 {
				dp[i][j] = 0
			} else if j == 1 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[rows][cols]
}

func MinPathSum(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	if rows == 0 || cols == 0 {
		return 0
	}
	dp := make([][]int, 0)
	for i := 0; i <= rows; i++ {
		dp = append(dp, make([]int, cols+1))
	}
	for i := 0; i <= cols; i++ {
		dp[0][i] = -1
	}
	for j := 0; j <= rows; j++ {
		dp[j][0] = -1
	}
	dp[1][1] = grid[0][0]
	for i := 2; i <= cols; i++ {
		dp[1][i] = dp[1][i-1] + grid[0][i-1]
	}
	return doMinPathSum(grid, dp, rows, cols)
}

func doMinPathSum(grid, dp [][]int, rows, cols int) int {
	for i := 2; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if j == 1 {
				dp[i][j] = grid[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = grid[i-1][j-1] + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[rows][cols]
}

func min(a, b int) int {
	if a == -1 {
		return b
	} else if b == -1 {
		return a
	} else if a <= b {
		return a
	} else {
		return b
	}
}

// 91 Decode Ways
func NumDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if 0 < s[0]-'0' && s[0]-'0' < 10 {
		dp[1] = 1
	} else {
		dp[1] = 0
	}
	for i := 2; i <= len(s); i++ {
		if 1 <= s[i-1]-'0' && s[i-1]-'0' <= 9 {
			dp[i] += dp[i-1]
		}
		front := s[i-2] - '0'
		if 10 <= front*10+s[i-1]-'0' && front*10+s[i-1]-'0' <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

// 95 Unique Binary Search Trees II
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTrees(n int) []*TreeNode {

	dp := make([][]*TreeNode, n+1)
	dp[0] = nil
	nodes := make([]*TreeNode, 0)
	node := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	dp[1] = append(nodes, node)
	return nil
}

// 96  Unique Binary Search Trees
func NumTrees(n int) int {
	if n < 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[i-j] * dp[j-1]
		}
	}
	return dp[n]
}
