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
