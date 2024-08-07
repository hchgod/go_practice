package main

import (
	"fmt"
	"sort"
)

type KeyValue struct {
	Key   string
	Value int
}

type ByValue []KeyValue

func (bv ByValue) Len() int           { return len(bv) }
func (bv ByValue) Less(i, j int) bool { return bv[i].Value < bv[j].Value }
func (bv ByValue) Swap(i, j int)      { bv[i], bv[j] = bv[j], bv[i] }

func main() {
	// 声明一个哈希表
	hashTable := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 20,
	}

	// 将键值对存入结构体切片
	var keyValueSlice []KeyValue
	for key, value := range hashTable {
		keyValueSlice = append(keyValueSlice, KeyValue{Key: key, Value: value})
	}

	// 使用 sort.Sort() 进行值的排序
	sort.Sort(ByValue(keyValueSlice))

	// 打印排序后的结果
	for _, kv := range keyValueSlice {
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}
}
