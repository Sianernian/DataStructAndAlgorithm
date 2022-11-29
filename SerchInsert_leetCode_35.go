package main

import "fmt"

/*
TODO: 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引.
	 如果目标值不存在于数组中，返回它将会被按顺序插入的位置,请必须使用时间复杂度为 O(log n) 的算法。
*/
func main() {
	a := []int{1, 3, 5, 6, 6, 78, 98, 546}
	fmt.Println(serchInsert(a, 9))
}

func serchInsert(nums []int, target int) int {
	low := 0
	a := len(nums)
	height := len(nums)

	for low < height {
		mid := (low + height) / 2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			a = mid
			height = mid
		} else {
			return mid
		}
	}
	return a
}
