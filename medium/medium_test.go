package medium

import (
	"fmt"
	"reflect"
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

func TestFindDuplicate(t *testing.T) {
	input := []int{1, 3, 4, 2, 2}
	input1 := []int{3, 1, 3, 4, 2}
	fmt.Println(FindDuplicate(input))
	fmt.Println(FindDuplicate(input1))
}

func TestUniquePaths(t *testing.T) {
	fmt.Println(UniquePaths(3, 2))
	fmt.Println(UniquePaths(7, 3))
	fmt.Println(UniquePaths(2, 1))
	fmt.Println(UniquePaths(1, 2))
}

func TestUniquePathsWithObstacles(t *testing.T) {
	input := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(UniquePathsWithObstacles(input))
	input1 := [][]int{
		{1},
	}
	fmt.Println(UniquePathsWithObstacles(input1))
}

func TestMinPathSum(t *testing.T) {
	//input := [][]int{
	//	{1, 3, 1},
	//	{1, 5, 1},
	//	{4, 2, 1},
	//}
	//fmt.Println(MinPathSum(input))
	s := "123456789"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d, %s, %v, %v\n", s[i]-'0', reflect.TypeOf(s[i]-'0'), s[i]-'0' > 0, s[i]-'0' < 26)
	}
}

func TestNumDecodings(t *testing.T) {
	fmt.Println(NumDecodings("12"))
	fmt.Println(NumDecodings("226"))
	fmt.Println(NumDecodings("111"))
	fmt.Println(NumDecodings("10"))
	fmt.Println(NumDecodings("01"))
	fmt.Println(NumDecodings("012"))
	fmt.Println(NumDecodings("12120"))
	fmt.Println(NumDecodings("1212"))
}

func TestMakeChild(t *testing.T) {
	node := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node1 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	node.Right = node1
	pNode := node
	node.Val = 5
	fmt.Println(pNode.Val)
	fmt.Println(pNode.Right)
	fmt.Println(node)
	node = node.Right
	fmt.Println(node)
	fmt.Println(pNode)
}

func TestNumTrees(t *testing.T) {
	fmt.Println(NumTrees(1))
	fmt.Println(NumTrees(2))
	fmt.Println(NumTrees(3))
	fmt.Println(NumTrees(4))
	fmt.Println(NumTrees(5))
}
