package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const kmToMars int = 62_100_000

func main() {
	rand.Seed(time.Now().Unix())
	spaceLines := []string{"Space Adventures", "Space X", "Virgin Galactic"}
	tripTypes := []string{"One-way", "Round-trip"}
	fmt.Printf("%-20s %4s %-10s  %5s\n", "Spaceline", "Days", "Trip Type", "Price")
	fmt.Printf("%45s\n", strings.Repeat("=", 45))
	for i := 0; i < 10; i++ {
		spaceLine := getSpaceLine(spaceLines)
		tripType := getTripType(tripTypes)
		days, price := calculateDayPrice(tripType)
		fmt.Printf("%-20s %4d %-10s  $%4d\n", spaceLine, days, tripType, price)
	}

}

func getSpaceLine(spaceLines []string) string {
	return spaceLines[rand.Intn(3)]
}

func calculateDayPrice(tripType string) (int, int) {
	randomInt := rand.Intn(15)
	kmPerSec := 16 + randomInt
	price := 0
	if tripType == "One-way" {
		price = 36 + randomInt
	} else {
		price = 2 * (36 + randomInt)
	}
	seconds := kmToMars / kmPerSec
	return seconds / (3600 * 24), price
}

func getTripType(tripTypes []string) string {
	return tripTypes[(rand.Intn(2))]
}
