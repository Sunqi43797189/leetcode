package leetcode

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	arr := []int{3,2,4}
	target := 6
	result := twoSum(arr, target)
	fmt.Println(result)
}
