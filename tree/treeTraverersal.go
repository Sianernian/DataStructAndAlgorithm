package main

// TODO 前序遍历 中左右 递归
func preTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node ==nil{   // 深度遍历 当没有节点时，跳出递归
			return
		}
		res =append(res,node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)

	return res
}

// TODO Morris遍历
func preMorrisTraversal(root *TreeNode)(res []int){
	var p1 =root
	for p1 !=nil{
		p2 :=p1.Left
		if p2 !=nil{
			for p2.Right !=nil &&  p2.Right !=p1{
				p2 =p2.Right
			}
			if p2.Right ==nil{
				res =append(res,p1.Val)
				p2.Right =p1
				p1 =p1.Left
				continue
			}
			p2.Right =nil
		}else {
			res =append(res,p1.Val)
		}
		p1 =p1.Right
	}
	return
}


// TODO 中序遍历  左中右
func midTraversal(root *TreeNode)(res []int){
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node ==nil{
			return
		}
		traversal(node.Left)
		res =append(res,node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

// TODO 后序遍历  左右中
func behindTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal= func(node *TreeNode) {
		if node==nil{
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res =append(res,node.Val)
	}
	traversal(root)
	return res
}