package main

import (
	"fmt"
)

func main() {
	//		       0  1   2  3    4
	arr := [5]int{5, 66, 7, 100, 1}

	fmt.Println(arr[1])

	arr[1] += 5

	fmt.Println("0:", arr[0])
	fmt.Println("1:", arr[1])
	fmt.Println("2:", arr[2])
	fmt.Println("3:", arr[3])
	fmt.Println("4:", arr[4])

	arr2 := [6]int{}
	arr2[0] = arr[0]
	arr2[1] = arr[1]
	arr2[2] = arr[2]
	arr2[3] = arr[3]
	arr2[4] = arr[4]
	arr2[5] = 77

	fmt.Println("arr", arr)
	fmt.Println("arr2", arr2)
}
