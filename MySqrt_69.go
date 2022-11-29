package main

import (
	"fmt"
)

/*
	TODO 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
	 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
	 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
*/
func main() {

	a:=mySqrt(16)
	fmt.Println(a)
}

func mySqrt(x int)int{
	if x ==0{
		return 0
	}
	low :=0
	height :=x
	for low< height {
		mid :=(low + height) /2
		if mid < x /mid {
			low = mid  +1
		}else if mid > x /mid{
			height =mid
		}else if mid ==x/mid{
			return mid
		}
	}
	
	return height -1
}
