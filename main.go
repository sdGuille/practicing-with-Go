package main

import "fmt"

// Array

func main() {
	myBand := [4]string{"Guillermo", "Katherine", "Diego", "Carlos"}
	arrayNumber := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	starCalification := [...]string{ // those three dots means undefined long of elements.
		"☆",
		"☆☆",
		"☆☆☆",
		"★★★★",
		"★★★★★",
	}

	currency := [...]string{0: "Canadian Dollar", 1: "Peso Mexicano", 2: "Dollar", 4: "Yuang"}

	fmt.Println(arrayNumber)
	fmt.Println(myBand)
	fmt.Println(starCalification)

	fmt.Println(currency[3]) // we don't
	fmt.Println(currency[2])
}
