package main

import "fmt"

// **** Slice **** //

func main() {
	numbers := []int{1, 2, 3, 4}

	numbers = append(numbers, 5)
	fmt.Println(numbers)

	newSlice := numbers[0:3]
	numbers[0] = 30

	fmt.Println(newSlice) // here we got 3 values from numbers variable
	fmt.Println(numbers)
	// slices are only a reference of a pice value from Array

	// let's see another example about slice

	months := []string{"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November"}

	// Pointer
	// Lenght
	// Capacity

	lenght := len(months)
	capacity := cap(months)

	// Printf
	fmt.Printf("The lenght is: %v, The capacity is: %v, direction %p \n", lenght, capacity, months)
	months = append(months, "December") // if the structure is on max capacity, this will generate new array

	lenght = len(months)
	capacity = cap(months)

	fmt.Printf("The lenght is: %v, The capacity is: %v, direction %p \n", lenght, capacity, months)
	// if we can see on the terminal the direction changes
	// this means: We are working with a new object

}
