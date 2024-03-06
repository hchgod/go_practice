package main

import (
	"fmt"
	"sort"
)

func main() {
	sliceint := []int{7, 5, 8, -2, 9, 1, 4}
	sort.Ints(sliceint)
	fmt.Println(sliceint)

	sort.Sort(sort.Reverse(sort.IntSlice(sliceint)))
	fmt.Println(sliceint)

	sliceStr := []string{"python", "java", "c++", "c", "golang", "PHP"}
	sort.Strings(sliceStr)
	fmt.Println(sliceStr)

	sort.Sort(sort.Reverse(sort.StringSlice(sliceStr)))
	fmt.Println(sliceStr)

}
