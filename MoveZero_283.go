package main


/*
	TODO 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
		  请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
*/
func main(){
	a:=[]int{0,1,0,3,12}
	moveZero(a)

}

func moveZero(nums []int){
	length :=len(nums)
	var rel []int
	result:=0
	for i := 0; i < length; i++ {
		if nums[i] !=0{
			nums[result] =nums[i]
			result ++
		}else if nums[i]==0{
			rel =append(rel,0)
		}
	}
	//nums =append(nums[:result],rel...)
	nums =append(nums[:0:0],nums[:result]...)
	nums=append(nums,rel...)

}
