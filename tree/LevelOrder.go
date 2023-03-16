package main

import "fmt"

func main() {
	a :=TreeNode{}
	a.setVal(3) // 3 920  157
	a.Left=&TreeNode{9,nil,nil}
	a.Right=&TreeNode{Val: 20}
	a.Right.Right=&TreeNode{Val: 15}
	a.Right.Left=&TreeNode{Val: 7}

	fmt.Println(levelOrder2(&a))
}


func levelOrder (root *TreeNode)[][]int{
	// 通过队列的方式进行操作
	que :=make([]*TreeNode,0)
	que =append(que,root)

	res :=make([][]int,0)
	if root ==nil{  // 如果为空，则 返回 空值
		return res
	}

	for len(que) !=0{
		res2 :=make([]int,0)  // 存放当前 元素
		size :=len(que)
		for i := 0; i < size; i++ {  // 不能直接 循环 len(que） 因为在循环加入了新的节点，使对列长度发生变化，会影响内层循环  遍历的就不是完整的队列
			fir :=que[0]
			res2 =append(res2,fir.Val)
			que =que[1:] // 弹出第一个值
			if fir.Left !=nil{ //拿左边值
				que =append(que,fir.Left)
			}
			if fir.Right !=nil{ // 右边值
				que=append(que,fir.Right)
			}
		}
		res =append(res,res2)
	}
	return res
}

func levelOrder2(root *TreeNode) [][]int {
	// 通过队列的方式进行操作
	que := make([]*TreeNode, 0)
	que = append(que, root)

	res := make([][]int, 0)
	if root == nil { // 如果为空，则 返回 空值
		return res
	}

	for len(que) != 0 {
		res2 := make([]int, 0) // 存放当前 元素
		for i := 0; i < len(que); i++ {
			fir := que[0]
			res2 = append(res2, fir.Val)
			que = que[1:] // 弹出第一个值
			if fir.Left != nil { //拿左边值
				que = append(que, fir.Left)
			}
			if fir.Right != nil { // 右边值
				que = append(que, fir.Right)
			}
		}
		res = append(res, res2)
	}
	return res
}