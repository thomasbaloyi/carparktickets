package main

import (
	"fmt"
)

func main() {

	// Tariff table

	type Tariff struct {
		Cost     byte
		Duration int64
	}

	// var tariffs = []Tariff{
	// 	{Cost: 10, Duration: 1000},
	// 	{Cost: 20, Duration: 3000},
	// 	{Cost: 30, Duration: 6000},
	// 	{Cost: 50, Duration: 8000},
	// 	{Cost: 90, Duration: 9000},
	// 	{Cost: 80, Duration: 10000},
	// }

	// Totally new to me, Go does not have a while loop, can be simulated as follows:
	for {
		var userInput = requestUserInput()
		if userInput == "q" {
			break
		}

		switch userInput {
		case "a":
			fmt.Println("Advancing clock...")
		case "l":
			fmt.Println("Leave carpark")
		case "e":
			fmt.Println("Enter carpark")
		}
	}

	// TODO: Advance a clock by x amount of time from user input using "a" or "advance"

	// TODO: Arrive at a car park using "arrive", calculate fair based on duration and tariff table

	// TODO: Depart car park
}

func requestUserInput() string {
	var userInput string
	fmt.Println("\n---------- Commands ----------")
	fmt.Println("(q) to quit \n(a) to advance clock \n(l) to leave car park \n(e) to enter car park")
	fmt.Print("Enter a command: ")
	fmt.Scan(&userInput)
	return userInput
}
