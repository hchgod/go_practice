package main

import (
	"fmt"
)

type errNotFound struct {
    file string
}

func (e *errNotFound) Error() string {              // 他是继承了  Error接口的结构体  自行百度
    return fmt.Sprintf("file %q not found", e.file)
}

func open(file string) error { //当return的时候被当做error类型时，自动的转成error类型  并且调用Error方法
    return &errNotFound{file: file}
}


func main() {
    if err := open("hello.txt"); err != nil {
        if _, ok := err.(*errNotFound); ok {
            fmt.Println("File not found error:", err)
        } else {
            panic("unknown error")
        }
    }
}
