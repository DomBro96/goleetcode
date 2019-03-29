package easy

func IntToRoman(num int) string {
	/**
			I V  X  L  C   D   M
	        1 5 10 50 100 500 1000
		    IV 4  IX 9
			XL 40 XC 90
			CD 400 CM 900
	*/
	result := ""
	var mCount, cCount, xCount, iCount int
	// 判断 1000
	for mCount = 0; mCount < num/1000; mCount++ {
		result += "M"
	}
	num -= mCount * 1000
	if num >= 900 { // 判断 900
		result += "CM"
		num -= 900
	} else if num >= 500 { // 判断 500
		result += "D"
		num -= 500
	} else if num >= 400 { // 判断 400
		result += "CD"
		num -= 400
	}

	for cCount = 0; cCount < num/100; cCount++ { // 判断 100
		result += "C"
	}
	num -= cCount * 100
	if num >= 90 { // 判断 90
		result += "XC"
		num -= 90
	} else if num >= 50 { //判断 50
		result += "L"
		num -= 50
	} else if num >= 40 {
		result += "XL"
		num -= 40
	}

	for xCount = 0; xCount < num/10; xCount++ {
		result += "X"
	}
	num -= xCount * 10

	if num >= 9 {
		result += "IX"
		num -= 9
	} else if num >= 5 {
		result += "V"
		num -= 5
	} else if num >= 4 {
		result += "IV"
		num -= 4
	}
	for iCount = 0; iCount < num/1; iCount++ {
		result += "I"
	}
	return result
}

/**
		罗马数字变阿拉伯数字
		I V  X  L  C   D   M
        1 5 10 50 100 500 1000
	    IV 4  IX 9
		XL 40 XC 90
		CD 400 CM 900
*/
func RomanToInt(s string) int {
	result := 0
	index := 0
	for index < len(s) {
		if s[index] == 'M' {
			result += 1000
			index++
			continue
		}
		if index+1 < len(s) && s[index] == 'C' && s[index+1] == 'M' {
			result += 900
			index += 2
			continue
		}
		if s[index] == 'D' {
			result += 500
			index++
			continue
		}
		if index+1 < len(s) && s[index] == 'C' && s[index+1] == 'D' {
			result += 400
			index += 2
			continue
		}
		if s[index] == 'C' {
			result += 100
			index++
			continue
		}
		if index+1 < len(s) && s[index] == 'X' && s[index+1] == 'C' {
			result += 90
			index += 2
			continue
		}
		if s[index] == 'L' {
			result += 50
			index++
			continue
		}
		if index+1 < len(s) && index+1 < len(s) && s[index] == 'X' && s[index+1] == 'L' {
			result += 40
			index += 2
			continue
		}
		if s[index] == 'X' {
			result += 10
			index++
			continue
		}
		if index+1 < len(s) && s[index] == 'I' && s[index+1] == 'X' {
			result += 9
			index += 2
			continue
		}
		if s[index] == 'V' {
			result += 5
			index++
			continue
		}
		if index+1 < len(s) && s[index] == 'I' && s[index+1] == 'V' {
			result += 4
			index += 2
			continue
		}
		if s[index] == 'I' {
			result += 1
			index++
			continue
		}
	}
	return result
}

// 35. Search Insert Position
func SearchInsert(nums []int, target int) int {
	if len(nums) < 1 {
		return 0
	}
	nums = append(nums, target)
	j := len(nums) - 1
	for j > 0 && nums[j-1] > target {
		j--
	}
	nums[j] = target
	if j > 0 && nums[j-1] == target {
		return j - 1
	}
	return j
}

// 202. happy number
func Ishappy(n int) bool {
	numMap := make(map[int]bool)
	num := sum(n)
	for num != 1 {
		if _, exists := numMap[num]; !exists {
			numMap[num] = true
			num = sum(num)
		} else {
			break
		}
	}

	if num == 1 {
		return true
	} else {
		return false
	}
}

func sum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

// 268. Missing Number
func MissingNumber(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	start := 0
	end := len(nums)
	result := 0
	// binary search
	for {
		if start == end {
			result = start
			break
		}
		middle := (start + end) / 2
		if countNum(nums, start, middle) == middle-start+1 {
			start = middle + 1
		} else {
			end = middle
		}

	}
	return result
}

// count range start to end
func countNum(nums []int, start, end int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= start && nums[i] <= end {
			count++
		}
	}
	return count
}

// 70
func ClimbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
