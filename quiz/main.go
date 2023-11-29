package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "problems.csv", "File name of file with problems in CSV format")
	flag.Parse()
	fmt.Printf("File name: %s\n", fileName)

	problems, err := ReadQuiz(fileName)
	if err != nil {
		log.Fatalf("%#v", err)
	}

	ShuffleQuiz(problems)
	numProblems, goodAnswers, err := PlayQuiz(problems)

	if err != nil {
		fmt.Printf("\nTimeout: you have solved %d problems out of %d\n", goodAnswers, numProblems)
	} else {
		fmt.Printf("You have solved %d problems out of %d\n", goodAnswers, numProblems)
	}
}

func ReadQuiz(s string) (problems [][]string, err error) {
	f, err := os.Open(s)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(f)
	return reader.ReadAll()
}

func ShuffleQuiz(problems [][]string) {
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})
}

func PlayQuiz(problems [][]string) (numProblems int, solvedProblems int, err error) {
	StartQuiz()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	numProblems = len(problems)
	solvedProblems = 0

	ch := make(chan int)
	go func() {
		for _, problem := range problems {
			question := problem[0]
			answer := problem[1]
			givenAnswer, _ := GetAnswer(question)
			if answer == givenAnswer {
				ch <- 1
			}
		}
	}()

	for i := 0; i < len(problems); i++ {
		select {
		case <-ch:
			solvedProblems++
		case <-ctx.Done():
			return numProblems, solvedProblems, ctx.Err()
		}
	}
	return numProblems, solvedProblems, nil
}

func GetAnswer(question string) (answer string, err error) {
	fmt.Printf("%s = ", question)
	reader := bufio.NewReader(os.Stdin)
	answer, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	answer = strings.TrimSuffix(answer, "\n")
	return answer, nil
}

func StartQuiz() {
	fmt.Print("You have 10 seconds to finish the quiz. Enter to start quiz: ")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	return
}
