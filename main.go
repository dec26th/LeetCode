package main

import "fmt"

func main() {
	slice := []int {1, 2, 3, 4, 5, 6}
	slice2 := slice[1:4]
	slice2 = append(slice2, 1)  // 1, 2, 3, 4, 1, 6
	slice1 := slice[1:2:2]
	fmt.Println(len(slice1), cap(slice1))
	slice1 = append(slice1, 1) //1, 2, 3, 4, 1, 6       2, 1
	fmt.Println(slice1)
	fmt.Println(slice)
}
