package main

/*
	 TODO 双指针  list 翻转
*/
func reverseList(head *Node) *Node {
	var pre *Node
	cur :=head
	for cur !=nil{
		a:=cur.Next
		cur.Next,pre =pre ,cur
		cur =a
	}
		return pre
}