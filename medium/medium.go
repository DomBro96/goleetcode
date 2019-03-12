package medium

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
