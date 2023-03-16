package main

import "fmt"

type TreeNode struct {
	Val	int
	Left *TreeNode
	Right *TreeNode
}


func (node *TreeNode)print(){
	if node ==nil{
		fmt.Println("nil=node")
	}else {
		fmt.Println(node.Val)
	}
}


func (node *TreeNode)setVal(val int){
	node.Val =val
}

func (node *TreeNode)recu(){
	if node !=nil{
		node.print()
		node.Left.recu()
		node.Right.recu()
	}else {
		return
	}
}