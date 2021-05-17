package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetName() string {
	fmt.Print("Your name = ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString(byte('\n'))
	if err != nil {
		log.Fatalln("Could not read your name")
	}
	return strings.TrimSpace(name)
}

func GetAge() int {
	fmt.Print("In what year are your born? ")
	reader := bufio.NewReader(os.Stdin)
	year, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Could not read the year you are born")
	}
	year = strings.TrimSpace(year)
	birthYear, err := strconv.Atoi(year)
	if err != nil {
		log.Fatalf("Could not convert string %s to integer", year)
	}
	currentYear := time.Now().Year()
	return currentYear - birthYear
}
