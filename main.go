package main

import "fmt"

// Read the Keyboard input with fmt.Scan

var (
	name   string
	age    int
	height float32
)

func main() {
	fmt.Print("Please write your name here: ")
	fmt.Scanf("%s", &name)
	fmt.Print("how old are you? ")
	fmt.Scanf("%d", &age)
	fmt.Print("Write your height: ")
	fmt.Scanf("%f", &height)
	fmt.Printf("Hi mi name is %s i'm %d years old, and my height is %.2f", name, age, height)

}
