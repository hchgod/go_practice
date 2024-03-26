// 
package main

import "fmt"

type EmptyInterface interface{}

func main() {
	contacts := make(map[string]EmptyInterface)
	contacts["Dot"] = int64(13111111111)
	contacts["Dukuan"] = "010-11111111"

	for k, v := range contacts {
		fmt.Printf("%s的联系方式是: %v\n", k, v)
	}
}