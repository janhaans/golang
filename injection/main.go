package main

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type FileLogger struct{}

func (l FileLogger) Log(message string) {
	fmt.Printf("[File Logger]: %s\n", message)
}

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Printf("[Console Logger]: %s\n", message)
}

type Application struct {
	logger Logger
}

func (a Application) Run() {
	a.logger.Log("Application is running")
}

func NewApplication(logger Logger) *Application {
	return &Application{logger}
}

func main() {
	logger := ConsoleLogger{}
	application := NewApplication(logger)
	application.Run()
}
