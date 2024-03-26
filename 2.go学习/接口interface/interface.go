package main

import (
	"fmt"
)
type dbcommon interface{
	insert(string) error
	delete(string) error 
}


type DBconfig struct {
	user string
	password string
	host string
	database string
}

type mysql struct {
	config DBconfig
	chatset string
}

type sql_server struct {
	config DBconfig
	chatset string
}

func (mys *mysql)insert(data string) (error) {
	fmt.Println("这是插入一个数据",data)
	return nil
}

func (mys *mysql)delete(data string) (error) {
	fmt.Println("这是删除一个数据",data)
	return nil
}

func (ss *sql_server)insert(data string) (error) {
	fmt.Println("这是插入一个数据",data)
	return nil
}

func (ss *sql_server)delete(data string) (error) {
	fmt.Println("这是删除一个数据",data)
	return nil
}


func main() {
	dataconfig := DBconfig{"root", "password", "127.0.0.1", "user"}
	dbtype := "mysql"
	var dbinterface dbcommon
	if dbtype == "mysql" {
		connection_db := mysql{
			chatset: "xxx",
			config: dataconfig,
		}
		dbinterface = &connection_db      //值类型传递给引用类型    结构体是值类型  
		// 切片（slice）、映射（map）、通道（channel）、接口（interface）、指针（pointer）
		//值类型包括基本数据类型（如int、float、bool、string）和结构体（struct）
	}else if dbtype == "sql_server" {
		connection_db := sql_server{
			chatset: "xxx",
			config: dataconfig,
		}
		dbinterface = &connection_db
	}	
	dbinterface.insert("3")
}