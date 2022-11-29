package main

type ListNode struct{
	Value	int
	Next	*ListNode
}

/*
	TODO 给你一个链表的头节点 head 和一个整数 val
		 请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
*/
func main(){
	//a:=&ListNode{}
	//fmt.Println(removeElement_list(a,1))

}

// 虚拟头节点   可以比不使用少一个 for循环 ， 删除了头节点，但next还是和target一样还得删，就得使用循环判断
func removeElement_list(head *ListNode,val int)*ListNode{
	dummyHead :=&ListNode{}
	dummyHead.Next =head
	e := dummyHead  // 临时指针 指向 虚拟头节点   为什么不指向 next  因为要删除的就是 next
	for e!=nil && e.Next !=nil{
		if e.Next.Value == val{
			e.Next =e.Next.Next
		}else {
			e = e.Next
		}
	}
	return dummyHead.Next  // 返回 head,head可能会被删除， 真正的head是 dummyHead next
}