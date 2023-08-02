package main

import (
	"fmt"
	"github.com/songjuncai1122/geekjob/slice"
)

func main() {
	i := []int{1, 3, 5, 7, 8, 9}
	inew := slice.Delete(i, 5)
	inew2 := slice.Delete2(i, 5)
	fmt.Println(inew)
	fmt.Println(inew2)
}
