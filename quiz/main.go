package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

	numProblems, goodAnswers := PlayQuiz(problems)

	fmt.Printf("You have solved %d problems out of %d\n", goodAnswers, numProblems)
}

func ReadQuiz(s string) (problems [][]string, err error) {
	f, err := os.Open(s)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(f)
	return reader.ReadAll()
}

func PlayQuiz(problems [][]string) (numProblems int, solvedproblems int) {
	goodAnswers := 0
	for _, problem := range problems {
		question := problem[0]
		answer := problem[1]
		givenAnswer, _ := GetAnswer(question)
		if answer == givenAnswer {
			goodAnswers++
		}
	}
	return len(problems), goodAnswers
}

func GetAnswer(question string) (string, error) {
	fmt.Printf("Question: %s\n", question)
	fmt.Print("Enter answer: ")
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	answer = strings.TrimSuffix(answer, "\n")
	return answer, nil
}
