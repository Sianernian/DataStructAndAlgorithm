package main

import "fmt"

func main() {
	a:=[]int{3,2,2,3}
	fmt.Println(removeElement(a,3))
}


/*
*	TODO 快慢指针法 删除element  双指针
*/
func removeElement(nums []int,val int) int{
	length :=len(nums)
	result :=0
	for i := 0; i < length; i++ {
		if nums[i] !=val{
			nums[result] =nums[i]
			result ++
		}
	}
	nums =nums[:result]
	return result
}


// 直接破解
//直接双层for 循环遍历出要的值
func removeElement1(nums []int,val int) int{
	a :=len(nums)
	for i := 0; i < a; i++ {
		if nums[i] ==val{
			for j := i+1; j < a ; j++ {
				nums[j-1] =nums[j];
			}
			i -- // 下标往前移动一位
			a-- // 数组大小 -1
		}
	}
	nums =nums[:a]
	return  a
}
