package main

import "fmt"

/*
	TODO: 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。
	 请你找出给定目标值在数组中的开始位置和结束位置。
	 如果数组中不存在目标值 target，返回 [-1, -1]。
	 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。  在一个循环里 / 2 就是 O(logN)

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]

提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109
*/

func main() {
	nums :=[]int{}
	a:=serchRange(nums,2)
	fmt.Println(a)
}

func serchRange(nums []int,target int) []int {
	leftfunc :=func(low1,height1 int)int{
		a :=-1
		for low1 <= height1 {
			mid :=(low1 +height1) /2
			if target < nums[mid] {
				height1 = mid - 1
			} else if target > nums[mid] {
				low1 = mid + 1
			} else {
				height1, a = mid-1, mid

			}
		}
		return a
	}

	rightfunc :=func(low1,height1 int)int{
		a :=-1
		for low1 <= height1 {
			mid :=(low1 +height1) / 2
			if target < nums[mid] {
				height1 = mid - 1
			} else if target > nums[mid] {
				low1 = mid + 1
			} else {
				low1, a = mid+1, mid
			}
		}
		return a
	}

	lefe:=leftfunc(0,len(nums)-1)  // len 要减1  不然数组下表越界
	right :=rightfunc(0,len(nums)-1)

	return []int{lefe,right}
}