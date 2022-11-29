package main

import "fmt"

func main() {
	a:=[]int{-1,0,3,4,6,8,9,10}

	r:=serch(a,-1)
	fmt.Println(r)
}

func serch(nums []int,target int) int {
		low :=0
		height :=len(nums)
		for low <height{
			mid := (low + height) /2
			if nums[mid] < target{
				low = mid +1
			}else if nums[mid] > target{
				height = mid
			}else {
				return mid
			}
		}
		return -1
}