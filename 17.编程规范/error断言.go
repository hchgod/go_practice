package main

import (
	"fmt"
)

type errNotFound struct {
	file string
}

func (e *errNotFound) Error() string {
	return fmt.Sprintf("file %q not found", e.file)
}

func IsNotFoundError(err error) bool {    // 自定义的错误断言函数
	_, ok := err.(*errNotFound)
	return ok
}

func Open(file string) error {
	return &errNotFound{file: file}
}

func main() {
	if err := Open("foo"); err != nil {
		if IsNotFoundError(err) {
			fmt.Println("File not found error:", err)
			// Handle the file not found error
		} else {
			panic("unknown error")
		}
	}
}