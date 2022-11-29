package main

import (
	"fmt"
	"math"
	"sort"
)

/*
	TODO 给你一个按 非递减顺序 排序的整数数组 nums，
 		 返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：
	输入：nums = [-4,-1,0,3,10]
	输出：[0,1,9,16,100]
	解释：平方后，数组变为 [16,1,0,9,100]
		  排序后，数组变为 [0,1,9,16,100]
*/
func main() {
	b:=[]int{-4,-1,0,3,10}

	a:=sortedSquares(b)
	fmt.Println(a)

}

func sortedSquares(nums []int)[]int{
	for i := 0; i <len(nums) ; i++ {
		nums[i] = int(math.Pow(float64(nums[i]), 2))
	}
	//arr :=make([]int,len(nums))
	//for i, v := range nums {
	//	arr[i] =v * v
	//}
	sort.Slice(nums, func(i, j int) bool {
		return 	nums[i] < nums[j]
	})  //  运行速度比 ints 快
	//sort.Ints(arr)
	//return arr
	return nums
}