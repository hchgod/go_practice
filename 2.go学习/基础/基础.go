package main
//结构体嵌套
import "fmt"

type phone struct {
    nums  string
    region string
}

type person struct {
    name string
    age int
    iphone phone    //结构体嵌套
}
func main() {
    var p person 
    p = person{
        name: "Tom",
        age: 25,
    }
    p.iphone.nums = "17737649832"
    p.iphone.region = "China"
    fmt.Println(p)
}
