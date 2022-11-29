package main

import "fmt"

/*
	TODO 给定一个含有 n 个正整数的数组和一个正整数 target 。
		 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，
		 并返回其长度。如果不存在符合条件的子数组，返回 0 。
示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
*/
func main(){
	a :=[]int{1,1,1,1,1,1,1,1}
	fmt.Println(minSubArrayLen(10,a))
}

// TODO 滑动窗口(暴力破解的思路)   快慢指针
func minSubArrayLen(target int ,nums []int)int{
	i:=0  // 控制 起始位置
	sum :=0  // sum 总和
	length :=len(nums)
	result := length +1  //  取最大值
	for j := 0; j < len(nums); j++ {
		sum +=nums[j]
		for sum >=target{
			subLen := j - i + 1  // 滑动窗口的长度
			if subLen < result{
				result = subLen  // 不断更新  suml 最小值
			}
			sum -=nums[i]
			i++
		}
	}

	if result == length +1{
		return 0
	}
	return result
}
