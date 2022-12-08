package main

import (
	"fmt"
)

// **** Map **** //

func main() {
	days := make(map[int]string)

	days[0] = "Sunday"
	days[1] = "Monday"
	days[2] = "Tuesday"
	days[3] = "Wednesday"
	days[4] = "Thursday"
	days[5] = "Friday"
	days[6] = "Saturday"

	// how to delete a key

	delete(days, 1)

	fmt.Println(days)

	// something more complicated

	users := make(map[string][]int)

	users["Guillermo"] = []int{9, 10, 8, 10, 9}

	fmt.Println(users) // this gonna print Guillermo Califications
}
