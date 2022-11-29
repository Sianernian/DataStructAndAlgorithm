package main

import "fmt"

/*
	TODO 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，
		 且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
*/
func main() {
	fmt.Println(generateMatrix(3))

}

func generateMatrix(n int)[][]int{
	top ,bottom,left,right:=0,n-1,0,n-1
	a:= n*n
	num :=1
	matrix :=make([][]int,n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int,n)
	}
	for num <= a {
		for i := left; i <= right; i++ {
			matrix[top][i] =num
			num ++
		}
		top ++
		fmt.Println("top:",matrix)
		for i := top; i <= bottom ; i++ {
			matrix[i][right] =num
			num++
		}
		right --
		fmt.Println("right:",matrix)
		for i := right; i >= left; i-- {
			matrix[bottom][i] =num
			num ++
		}
		bottom --
		fmt.Println("bottom:",matrix)
		for i := bottom; i >=top ; i-- {
			matrix[i][left] =num
			num ++
		}
		left ++
		fmt.Println("left:",matrix)
	}

	return matrix
}