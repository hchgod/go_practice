package main
//  这是多接口  多结构体   接口类型的切片
import "fmt"

type Shaper interface {
    Area() float32
}
type Square struct {
    side float32
}
func (sq *Square) Area() float32 {
    return sq.side * sq.side
}
type Rectangle struct {
    length, width float32
}
func (r Rectangle) Area() float32 {
    return r.length * r.width
}
func main() {

    r := &Rectangle{5, 3} // Area() of Rectangle needs a value
    q := &Square{5}      // Area() of Square needs a pointer
    // shapes := []Shaper{Shaper(r), Shaper(q)}
    // or shorter
    shapes := []Shaper{r, q}
    fmt.Println("Looping through shapes for area ...")
    for n, face := range shapes {
        fmt.Println(n,"Shape details: ", face)
        fmt.Println("Area of this shape is: ", face.Area())
    }
}