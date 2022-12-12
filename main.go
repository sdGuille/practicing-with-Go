package main

import (
	"fmt"
)

// **** Control Flow Conditionals **** //

func main() {
	users := map[int]string{} // make
	users[1] = "diego"
	users[2] = "kathy"
	users[3] = "guillo"
	users[4] = "unknow"

	for key, value := range users {
		fmt.Println("User number: ", key, "= ", value)
	}
}
