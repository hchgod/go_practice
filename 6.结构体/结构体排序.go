package main

import (
    "log"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    persons := []Person{
       {Name: "hch", Age: 18},
       {Name: "lv", Age: 20},
       {Name: "hb", Age: 19},
       {Name: "zb", Age: 30},
       {Name: "gg", Age: 23},
    }

    sort.Slice(persons, func(i, j int) bool {
       return persons[i].Age > persons[j].Age
    })

    for _, value := range persons {
       log.Println(value.Name, value.Age)
    }
}