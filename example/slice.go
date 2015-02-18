package main

import "fmt"

func main() {
	sInts := []int{1, 2, 3, 4, 5}
	smallInts := sInts[:2]
	fmt.Println(sInts[:])
	fmt.Println(smallInts)
	sInts[1] = 10
	fmt.Println(smallInts)
}
