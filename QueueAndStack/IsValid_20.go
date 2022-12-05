package main

import "fmt"


/*
	TODO 有效的括号给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

示例 1：
输入：s = "()"
输出：true
*/
func main(){
	s :="{()]"
	fmt.Println(isValid(s))
}

/*
  三种情况
	一: 遍历string 分别把对应的括号（遇到左括号把右括号放进）放进 栈中，等到 遍历到栈中相同时弹出，栈中的括号有多
	二：遍历string 分别把对应的括号（遇到左括号把右括号放进）放进 栈中，遇到右括号和栈中的不匹配，不匹配
	三：遍历string 分别把对应的括号（遇到左括号把右括号放进）放进 栈中，等到 遍历到栈中相同时弹出，栈中的括号少了
可以使用 hashMap
*/
func isValid(s string)bool{
	m:=map[byte]byte{ // 因为从左往右遍历  要和 判断反着使用
		']':'[',
		'}':'{',
		')':'(',
	}
	result :=make([]byte,0)
	for i := 0; i < len(s); i++ {
		if s[i] =='(' || s[i] =='[' ||s[i] =='{' {
			result =append(result,s[i])
		}else if len(result) >0 && result[len(result)-1] ==m[s[i]]{
			result =result[:len(result)-1]
		}else {
			return false
		}
	}

	return len(result) == 0
}