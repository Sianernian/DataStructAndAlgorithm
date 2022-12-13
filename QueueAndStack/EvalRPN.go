package main

import (
	"strconv"
)

/*
	TODO 根据 逆波兰表示法，求表达式的值。
	 	有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
		注意 两个整数之间的除法只保留整数部分。
		可以保证给定的逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中
示例 1：
输入：tokens = ["2","1","+","3","*"]
输出：9
解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
*/
func evalRPN(tokens []string)int{
	RPNStack :=[]int{}
	for _, token := range tokens {
	val ,err:=strconv.Atoi(token)
	if err !=nil{
		num1 ,num2:=RPNStack[len(RPNStack)-2],RPNStack[len(RPNStack)-1]
		RPNStack =RPNStack[:len(RPNStack)-2]
		switch token {
		case "+":
			RPNStack =append(RPNStack,num1+num2)
		case "-":
			RPNStack =append(RPNStack,num1-num2)
		case "*":
			RPNStack =append(RPNStack,num1*num2)
		case "/":
			RPNStack =append(RPNStack,num1/num2)
		}
	}else {
		RPNStack =append(RPNStack,val)
	}
	}
	return RPNStack[0]
}

