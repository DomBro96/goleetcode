package medium

import (
	"fmt"
	"testing"
)

func TestBinSearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(BinSearch(nums, 3))
	fmt.Println(BinSearch(nums, 6))
}

func TestSearch(t *testing.T) {
	//nums := []int{3, 4, 5, 1, 2}
	//	//fmt.Println(Search(nums, 4))
	//	//fmt.Println(Search(nums, 2))
	//	//fmt.Println(Search(nums, 6))
	nums1 := []int{0}
	fmt.Println(Search(nums1, 1))
}

func TestSearchRange(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4, 5}
	fmt.Println(SearchRange(nums, 3))
	fmt.Println(SearchRange(nums, 9))
}

func TestIsValidSudoku(t *testing.T) {
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
	is := IsValidSudoku(input)
	fmt.Println(is)
}
