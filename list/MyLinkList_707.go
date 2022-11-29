package main

import (
	"errors"
	"fmt"
)

type MyLinkList struct{
	Len int
	//headNode *Node
	//Tail *Node
	DummyHead *Node
}

type Node struct{
	Value int
	Next  *Node
}


// 设计链表
func main() {
	a:=Constructor()
	a.AddAtHead(1)
	//a.AddAtHead(2)
	a.AddAtTail(3)
//	a.AddAtTail(4)
	//a.AddAtTail(5)
	a.AddAtIndex(1,2) // 1  2  3
	//fmt.Println(a.Get(1)) // 2
	//a.AddAtTail(7)//  2 1 0 3 4 5 6 7
//	a.AddAtIndex(4,6)
	a.DeleteAtIndex(0) // 2 3
	fmt.Println(a.Get(0))// 2
	//fmt.Println(a.Len)


	//fmt.Println(a.DummyHead)
	//for{
	//	if a.DummyHead.Next !=nil{
	//		a.DummyHead =a.DummyHead.Next
	//	}else {
	//		break
	//	}
	//	fmt.Print(a.DummyHead.Value)
	//}
}

/*
	TODO Constructor() 构造函数
		Get() 获取链表第 index 个节点值
		AddAtHead() 在链表头节点添加一个节点，使其成为头节点
		AddAtTail()	在链表为添加一个节点
		AddAtIndex()	在index位置添加一个节点
		DeleteAtIndex()	在第index位置删除一个节点
*/
func Constructor()MyLinkList{
	return MyLinkList{
		0,
		//nil,
		//nil,
		&Node{Next: nil},
	}
}

func (this *MyLinkList)Get(index int)int{
	if index < 0 || index >=this.Len{
		return -1
	}
	cur :=this.DummyHead
	for i := 0; i < index; i++ {
		cur =cur.Next
	}
	return cur.Value
}

func (this *MyLinkList)AddAtHead(val int){
	node :=&Node{Value: val}
	if this.Len ==0{  // 去除默认时的 0值
		this.DummyHead =node
		this.Len++
		return
	}
	node.Next = this.DummyHead
	this.DummyHead =node
	this.Len ++
}

func (this *MyLinkList)AddAtTail(val int){
	node :=&Node{Value: val}
	cur :=this.DummyHead
	for cur.Next !=nil {
		cur =cur.Next
	}
	cur.Next =node
	this.Len ++
}


func (this *MyLinkList)AddAtIndex(index int,val int){
	if index < 0 || index > this.Len {
		fmt.Println(errors.New("输入错误"))
		return
	}
	node :=&Node{Value: val}
	cur :=this.DummyHead
	for i := 1; i < index -1 ; i++ {
		cur =cur.Next
	}
	node.Next =cur.Next
	cur.Next =node
	this.Len ++
}

func (this *MyLinkList)DeleteAtIndex(index int){
	if index <0 ||index > this.Len {
		fmt.Println(errors.New("输入错误"))
		return
	}else if index ==0{ // 要删除的时头节点
		this.DummyHead =this.DummyHead.Next
	//	fmt.Println("dummy.Next:",this.DummyHead)
		return
	}
	cur :=this.DummyHead
	for i := 1 ; i < index-1  ; i++ {
		cur =cur.Next
	}
	cur.Next =cur.Next.Next
	this.Len --
}