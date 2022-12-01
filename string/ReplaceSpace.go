package main

import "fmt"

func main(){
	a :="Hello w o r l d"

	fmt.Println(replaceSpace(a))
}

func replaceSpace(s string)string{
	StrByte :=[]byte(s)
	result :=make([]byte,0)
	for i := 0; i < len(StrByte); i++ {
		if StrByte[i] == ' ' {
			result =append(result,[]byte("%20")...)
		}else {
			result =append(result,StrByte[i])
		}
	}
	return string(result)
}
