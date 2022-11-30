package main

// string 反转
/*
	输入：s = ["h","e","l","l","o"]
	输出：["o","l","l","e","h"]
*/
func reverseString(s []byte){
	left:=0
	right :=len(s) -1
	for left < right {
		s[left],s[right] =s[right],s[left]
		left ++
		right --
	}
}