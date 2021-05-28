package trip

import (
	"fmt"
	"math/rand"
	"time"
)

const kmToMars int = 62_100_000

type TripToMars struct {
	SpaceLine string
	Days      int
	TripType  string
	Price     int
}

func (t TripToMars) Format() string {
	return fmt.Sprintf("%-20s %4d %-10s  $%4d", t.SpaceLine, t.Days, t.TripType, t.Price)
}

func (t TripToMars) Print() {
	fmt.Println(t.Format())
}

func GetTripToMars(spaceLines []string, tripTypes []string) TripToMars {
	trip := TripToMars{}
	trip.SpaceLine = getSpaceLine(spaceLines)
	trip.TripType = getTripType(tripTypes)
	trip.Days, trip.Price = calculateDayPrice(trip.TripType)
	return trip
}

func getSpaceLine(spaceLines []string) string {
	rand.Seed(time.Now().UnixNano())
	return spaceLines[rand.Intn(3)]
}

func calculateDayPrice(tripType string) (int, int) {
	rand.Seed(time.Now().UnixNano())
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
