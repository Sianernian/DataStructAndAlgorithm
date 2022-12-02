package main

type MyStack struct {
	InQueue []int
	OutQueue []int
}

type MyStack1 struct {
	Queue []int
}

func main() {
	//a:=Constructor()
	//a.Push(1)
	//a.Push(2)
	//fmt.Println(a)
	//fmt.Println(a.Top())
	//fmt.Println(a.Top())
	//a.Pop()
	//a.Pop()
	//fmt.Println(a)
	//fmt.Println(a.Empty())

}

func Constructor()MyStack{
	return MyStack{
		make([]int,0),
		make([]int,0),
	}
}

func (this *MyStack)Push(x int) {
	this.InQueue = append(this.InQueue,x)
	for len(this.OutQueue) >0{// 把in 导入 out中
		this.InQueue =append(this.InQueue,this.OutQueue[0])
		this.OutQueue =this.OutQueue[1:]
	}
	this.OutQueue,this.InQueue =this.InQueue,this.OutQueue

}

func (this *MyStack) Pop() int {
	result :=this.OutQueue[0]
	this.OutQueue =this.OutQueue[1:]
	return result
}


func (this *MyStack)Top() int {
		return this.OutQueue[0]
}

func (this *MyStack)Empty() bool {
	return len(this.InQueue) ==0 && len(this.OutQueue) ==0
}

func (this *MyStack1)Push1(x int){
	this.Queue =append(this.Queue,x)
}
func (this *MyStack1)Pop1() int {
	length :=len(this.Queue)-1
	for length !=0{
		result :=this.Queue[0]
		this.Queue =this.Queue[1:]
		this.Queue =append(this.Queue,result)
		length--
	}

	result :=this.Queue[0]
	this.Queue =this.Queue[1:]
	return result
}
func (this *MyStack1)Top1() int {
	result:=this.Pop1()
	this.Queue =append(this.Queue,result)
	return result
}
