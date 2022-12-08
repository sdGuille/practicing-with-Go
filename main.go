package main

import "fmt"

// **** Make function Slice **** //

func main() {
	myNewSlice := make([]int, 3, 3)

	myNewSlice[0] = 100
	myNewSlice[1] = 200
	myNewSlice[2] = 300

	myNewSlice = append(myNewSlice, 400)

	fmt.Println(myNewSlice)
	fmt.Println(len(myNewSlice))
	fmt.Println(cap(myNewSlice))

}
