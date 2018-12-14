package easy

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	num := 3
	result := IntToRoman(num)
	fmt.Println(result)
	num = 4
	result = IntToRoman(num)
	fmt.Println(result)
	num = 58
	result = IntToRoman(num)
	fmt.Println(result)
	num = 1994
	result = IntToRoman(num)
	fmt.Println(result)
}

func TestRomanToInt(t *testing.T) {
	s0 := "III"
	fmt.Println(RomanToInt(s0))
	s1 := "IV"
	fmt.Println(RomanToInt(s1))
	s2 := "IX"
	fmt.Println(RomanToInt(s2))
	s3 := "LVIII"
	fmt.Println(RomanToInt(s3))
	s4 := "MCMXCIV"
	fmt.Println(RomanToInt(s4))
}

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	fmt.Println(SearchInsert(nums, 5))
	fmt.Println(SearchInsert(nums, 2))
	fmt.Println(SearchInsert(nums, 7))
	fmt.Println(SearchInsert(nums, 0))
}
