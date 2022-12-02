package main

import "fmt"

/*
	 TODO 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
	 	实现 MyQueue 类：
		1.void push(int x) 将元素 x 推到队列的末尾
		2.int pop() 从队列的开头移除并返回元素
		3.int peek() 返回队列开头的元素
		4.boolean empty() 如果队列为空，返回 true ；否则，返回 false
example 1
输入：
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 1, 1, false]
*/

//俩个栈 模仿队列
type MyQueue struct{
	inStack,outStack []int
}

func main(){
	a:=constructor()
	a.Push(1)
	a.Push(2)
	a.Push(3)
	a.Push(4)
	fmt.Println(a)
	a.Pop()
	fmt.Println(a)
	a.Push(5)
	fmt.Println(a)
	a.Pop()
	fmt.Println(a)
	a.Pop()
	fmt.Println(a)
	a.Pop()
	fmt.Println(a)
	a.Pop()
	fmt.Println(a)
}

func constructor()MyQueue{
	return MyQueue{
		inStack: make([]int,0),
		outStack: make([]int,0),
	}
}

func (this *MyQueue)Push(x int)  {
	// 可以先进行判断是否为空
	for len(this.inStack) !=0{
		result :=this.inStack[len(this.inStack)-1]
		this.inStack =this.inStack[:len(this.inStack)-1]
		this.outStack = append(this.outStack,result)
	}
	this.inStack =append(this.inStack,x)
}

// 判断是否为空，如果为空需要把in 全倒入out中
func (this *MyQueue)Pop()int{
	for len(this.inStack) !=0{ // 导入到out中
		result :=this.inStack[len(this.inStack)-1]
		this.inStack =this.inStack[:len(this.inStack)-1]
		this.outStack = append(this.outStack,result)
	}
	if len(this.outStack) ==0{
		return -1
	}

	result :=this.outStack[len(this.outStack)-1] // 弹出第一个值
	this.outStack =this.outStack[:len(this.outStack)-1]
	return result
}

func (this *MyQueue)Peek()int{
	result :=this.Pop()
	if result ==0{
		return 0
	}
	this.outStack =append(this.outStack,result)
	return result
}

func (this *MyQueue)Empty()bool{
	return len(this.inStack) ==0 && len(this.outStack)==0
}



