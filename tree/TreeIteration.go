package main

//  TODO 二叉树 迭代遍历
func preIter(root *TreeNode) []int{
	var res []int
	var stack []*TreeNode

	for len(stack) > 0 || root !=nil{
		for root !=nil {
			res =append(res,root.Val)
			stack =append(stack,root.Right) //右节点 入栈
			root =root.Left
		}

		index :=len(stack)-1  // 栈顶
		root =stack[index] // 出栈
		stack =stack[:index]

		//node :=stackLisk.Remove() //获取弹出元素
		//// 弹出元素
		//	res =append(res,node.val)
		//if node.right !=nil{  // 栈先进后出 先处理右节点
		//	stackLisk.PushFront(node.right)
		//}
		//if node.left !=nil{
		//	stackLisk.PushFront(node.left)
		//}
	}

	return res
}


func behindIter(root *TreeNode)  []int{
	var res []int
	stack :=[]*TreeNode{root}
	for len(stack)>0 || root !=nil {
		for root !=nil{
			stack =append(stack,root.Right)
			root =root.Left
			res =append(res,root.Val)
		}
		index :=len(stack) -1
		root =stack[index]
		stack =stack[:index]
	}
	return  res
}