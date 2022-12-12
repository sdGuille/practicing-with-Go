package main

import "fmt"

// **** Conditionals **** //

func main() {
	var age int
	fmt.Println(" Enter your age: ")
	fmt.Scanf("%d", &age)

	if age < 18 {
		fmt.Println("You are not able to continue")
	} else {
		fmt.Println("Welcome to this forum")
	}

	var calification int

	fmt.Print("Write your calification: ")
	fmt.Scanf("%d", &calification)

	if calification == 10 {
		fmt.Println("Awsome your calification is perfect!")
	} else if calification == 8 || calification == 9 {
		fmt.Println("Great job!")
	} else if calification == 6 || calification == 7 {
		fmt.Println("You pass the grade but need study more")
	} else if calification >= 0 && calification <= 5 {
		fmt.Println("You Fail the grade :(")
	} else {
		fmt.Println("Invalid Calification")
	}

}
