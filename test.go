package main

import (
	"fmt"
	"github.com/nint8835/boostertrack/pkg/availability"
)

func main() {
	fmt.Printf("%#+v\n", availability.Locations)
	availabilities, err := availability.GetAvailabilities()
	if err != nil {
		fmt.Printf("Error getting availabilities: %v\n", err)
	}
	fmt.Printf("%#+v\n", availabilities)
}
