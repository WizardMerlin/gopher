package main

import (
	"fmt"
)


func main() {

	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28

	days, ok := dayMonths["Jans"] 
	if !ok {
		fmt.Println("Can't get days for Jan")
	} else {
		fmt.Printf("%d\n", days)
	}

	for month, day := range dayMonths {
		fmt.Println(month, "has ", day, "days")
	}
}