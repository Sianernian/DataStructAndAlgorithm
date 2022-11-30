package main

import "fmt"

func main() {
	a := "你好吗，卧槽的"
	fmt.Println(reverseStr(a,2))
}

func reverseStr(s string,k int)string{
	str :=[]rune(s)
	length:=len(s)
	for i := 0; i < length; i += 2*k {
		if i+k <length {
			reverse(str[i:i+k])
		}else {
			reverse(str[i:length])
		}

	}
	return string(str)
}

func reverse(s []rune){
	left :=0
	right :=len(s) -1
	for left <right{
		s[left],s[right] = s[right],s[left]
		left ++
		right --
	}
}