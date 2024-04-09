package main

import (
	"fmt"
)

func combine(n int, k int) (res [][]int) {
	list := []int{}
	var backtracking func(n int, k int,start int)
	backtracking = func(n int, k int,start int) {
		if len(list) == k {
            temp := make([]int,len(list))
            copy(temp,list)
			res = append(res, temp)
			return 	
		}
		for i:=start; i<=n; i++ {
			if n-i+1<k-len(list) {
				break
			}
			list = append(list,i)
			backtracking(n,k,i+1)
			list = list[:len(list)-1]
		}
		return 
	}
    backtracking(n,k,1)
	return res 
}

func main() {
	res := combine(4,3)
	fmt.Println(res)
}
