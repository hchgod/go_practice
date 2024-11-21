package main

import (
	"fmt"
)

func lemonadeChange(ratings []int) bool {
	five := 0
	ten := 0
	for i := 0; i < len(ratings); i++ {
		if ratings[i] == 5 {
			five++
			continue
		}
		if ratings[i] == 10{
			if five > 0{
				five--
				ten++
				continue
			} else {
				return false
			}
		}
		if ratings[i] == 20 && five > 0{
			five--
			if ten > 0 {
				ten--
				continue
			} else if five >= 2{
				five = five - 2
				continue
			} else {
				return false
			}
		}else{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(lemonadeChange([]int{5,5,5,10,20}))
	fmt.Println(lemonadeChange([]int{5,5,10}))

}
