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
