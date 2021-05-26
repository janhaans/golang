package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/janhaans/tickettomars/triptomars"
)

func main() {
	rand.Seed(time.Now().Unix())
	spaceLines := []string{"Space Adventures", "Space X", "Virgin Galactic"}
	tripTypes := []string{"One-way", "Round-trip"}
	fmt.Printf("%-20s %4s %-10s  %5s\n", "Spaceline", "Days", "Trip Type", "Price")
	fmt.Printf("%45s\n", strings.Repeat("=", 45))
	for i := 0; i < 10; i++ {
		tripToMars := triptomars.GetTripToMars(spaceLines, tripTypes)
		fmt.Println(tripToMars.Format())
	}

}
