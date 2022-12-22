package main

import (
	"fmt"
	"os"
)

// **** Conditionals **** //

func main() {
	// fmt.Println("Save the World with GO!!!")
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

}
