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
            fmt.Println("这是list的地址",&list[0])
            fmt.Println("这是temp的地址",&temp[0])
			res = append(res, temp)
            fmt.Println("这是res的地址",&res[0][0])
            if len(res) == 3 {
                fmt.Println("这是res[2][0]的地址",&res[2][0])
            }
			return 	
		}
		for i:=start; i<=n; i++ {
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
	res := combine(4,2)
	fmt.Println(res)
}
