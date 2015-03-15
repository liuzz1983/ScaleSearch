package main

import (
	"fmt"
	"github.com/liuzz1983/ScaleSearch/core"
)

func main() {
	fmt.Println("hello api")

	index, _ := core.NewLsiIndex("zhong")
	fmt.Println(index)
}
