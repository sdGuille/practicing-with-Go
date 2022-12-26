package main

import (
	"fmt"
)

// **** Enums declaration **** //

type DayOfTheWeek uint8

const (
	Sunday DayOfTheWeek = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	// fmt.Println("Save the World with GO!!!")
	fmt.Printf("Monday is %d\n", Monday)
	fmt.Printf("Friday is %d\n", Friday)
	fmt.Printf("Sunday is %d\n", Sunday)

}
