package main

import (
	"fmt"
	"strings"

	"github.com/janhaans/tickettomars/trip"
)

func main() {
	spaceLines := []string{"Space Adventures", "Space X", "Virgin Galactic"}
	tripTypes := []string{"One-way", "Round-trip"}
	fmt.Printf("%-20s %4s %-10s  %5s\n", "Spaceline", "Days", "Trip Type", "Price")
	fmt.Printf("%45s\n", strings.Repeat("=", 45))
	for i := 0; i < 10; i++ {
		tripToMars := trip.GetTripToMars(spaceLines, tripTypes)
		tripToMars.Print()
	}
}
