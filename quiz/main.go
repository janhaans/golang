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

type Problem struct {
	Question string
	Answer   string
}

type Quiz struct {
	Problems []Problem
}

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "problems.csv", "File name of file with problems in CSV format")
	flag.Parse()
	fmt.Printf("File name: %s\n", fileName)

	quiz, err := ReadQuiz(fileName)
	if err != nil {
		log.Fatalf("%#v", err)
	}

	ShuffleQuiz(quiz)
	numProblems, goodAnswers, err := PlayQuiz(quiz)

	if err != nil {
		fmt.Printf("\nTimeout: you have solved %d problems out of %d\n", goodAnswers, numProblems)
	} else {
		fmt.Printf("You have solved %d problems out of %d\n", goodAnswers, numProblems)
	}
}

func ReadQuiz(s string) (Quiz, error) {
	f, err := os.Open(s)
	if err != nil {
		return Quiz{}, err
	}

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return Quiz{}, err
	}

	quiz := Quiz{}
	for _, row := range records {
		problem := Problem{row[0], row[1]}
		quiz.Problems = append(quiz.Problems, problem)
	}

	return quiz, nil
}

func ShuffleQuiz(quiz Quiz) {
	rand.Shuffle(len(quiz.Problems), func(i, j int) {
		quiz.Problems[i], quiz.Problems[j] = quiz.Problems[j], quiz.Problems[i]
	})
}

func PlayQuiz(quiz Quiz) (numProblems int, solvedProblems int, err error) {
	StartQuiz()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	numProblems = len(quiz.Problems)
	solvedProblems = 0

	ch := make(chan int)
	go func() {
		for _, problem := range quiz.Problems {
			answer, _ := GetAnswer(problem.Question)
			if answer == problem.Answer {
				ch <- 1
			}
		}
	}()

	for i := 0; i < len(quiz.Problems); i++ {
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
