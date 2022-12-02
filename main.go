package main

import "fmt"

/*
	go build main.go -> an excecutable file
 	go run  main.go -> this gone excecute file internaly
*/
// this can be more simple
// const Sunday int = 0
// const Monday int = 1
// const Tuesday int = 2
// const Wednesday = 3
// const Thursday = 4
// const Friday = 5
// const Saturday = 6

// const (
// 	Sunday    int = 0
// 	Monday    int = 1
// 	Tuesday   int = 2
// 	Wednesday     = 3
// 	Thursday      = 4
// 	Friday        = 5
// 	Saturday      = 6
// )
// but can be even better
const (
	Sunday int = iota // the default value is "0"
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Saturday)
}
