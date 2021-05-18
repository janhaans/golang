package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	targetNumber := rand.Intn(100) + 1

	reader := bufio.NewReader(os.Stdin)

	guessed := false

	for i := 1; i <= 10; i++ {
		fmt.Print("Guess the target number (1 to 100) ")
		strNumber, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Could not read the guessed number")
		}
		strNumber = strings.TrimSpace(strNumber)
		number, err := strconv.Atoi(strNumber)
		if err != nil {
			log.Fatalf("Could not convert string %s to int", strNumber)
		}

		if number == targetNumber {
			guessed = true
			break
		} else if number < targetNumber {
			fmt.Println("Your guessed number is too LOW")
		} else {
			fmt.Println("Your guessed number is HIGH")
		}
	}

	if guessed == true {
		fmt.Println("You have guessed the right number")
	} else {
		fmt.Println("You have not guessed the right number")
	}
}
