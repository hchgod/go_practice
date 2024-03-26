package main
array 的基本操作
import (
	"fmt"
)

func get_value(array []int,index int) {
	for i := range array {
		if i == index {
			fmt.Println(array[i])
		}		
	}
}

func insert_value(array [10]int, num int, index int) {
	for i := len(array) - 1; i >= index; i-- {
		array[i+1] = array[i]
	}
	array[index] = num
	fmt.Println(array)
}

func delete_value(arr []int, index int) {
	for i := index; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr = arr[:len(arr)-1]  
	fmt.Println(arr)
}


func find_value(arr []int, index int)  (flag bool){
	for i,elemet := range arr {
		fmt.Println(i,elemet)
	}
	return false

}

func extend(arr []int,extend_length int,) ([] int){
	result := make([]int,len(arr)+extend_length)
	for i,element := range arr{
		result[i] = element
	}
	fmt.Println(result)
	return result
}


func main() {
	// arr := []int{4, 995, 545, 346}
	// get_value(arr,3)
	// insert_value(arr, 546,2)
	// delete_value(arr, 3)
	// find_value(arr,4)
	// extend(arr,3)
}