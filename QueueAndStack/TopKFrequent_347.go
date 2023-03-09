package main

import (
	"fmt"
	"sort"
)

func main (){
	s1 := []int{1,4,4,2,2,2,2,4,3,3}
	m1 :=make(map[int]int)
	s2 :=make([]int,0)
	for _, v := range s1 {
			m1[v]++
	}
	for k := range m1 {
		s2 =append(s2,k)
	}
	fmt.Println(m1)
	fmt.Println(s2)
	sort.Slice(s2, func(i, j int) bool {
		return m1[s2[i]] >m1[s2[j]]  // 出现频率 重大到小排序
	})

	s2 =s2[:3]
	fmt.Println(s2)

}

/*
		TODO 给你一个整数数组 nums 和一个整数 k，请你返回其中出现频率前 k 高的元素。可以按任意顺序返回答案。
	示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
*/
func topKfrequent(nums []int,k int)[]int{
	m1 :=make(map[int]int)
	s2 :=make([]int,0)
	for _, v := range nums {
		fmt.Println(m1[v])
		m1[v]++
	}
	fmt.Println(m1)

	for k1 := range m1 {
		s2 =append(s2,k1)
	}

	sort.Slice(s2, func(i, j int) bool {
		return m1[s2[i]] >m1[s2[j]]
	})

	return s2[:k]
}