package main

import (
	"context"
	"fmt"
	"time"
)

func LongFunc(n int64) string {
	time.Sleep(time.Duration(n) * time.Millisecond)
	return "I'm awake"
}

func LongFuncWithTimeout(n int64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	ch := make(chan string)
	go func() {
		s := LongFunc(n)
		ch <- s
	}()

	select {
	case msg := <-ch:
		return msg, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	resp, err := LongFuncWithTimeout(10)
	if err != nil {
		fmt.Println("Timeout!")
	}
	fmt.Println(resp)

	resp, err = LongFuncWithTimeout(30)
	if err != nil {
		fmt.Printf("Timeout!")
	}
	fmt.Println(resp)

	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch2 <- 2
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	select {
	case msg1 := <-ch1:
		fmt.Printf("Received message = %d from channel ch1\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received message = %d from channel ch2\n", msg2)
	//case <-time.After(5 * time.Millisecond):
	case <-ctx.Done():
		fmt.Println("Timeout!!!")
	}

}
