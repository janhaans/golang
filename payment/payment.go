package main

import (
	"fmt"
	"sync"
	"time"
)

type Payment struct {
	From   string
	To     string
	Amount float64
	once   sync.Once
}

func (p *Payment) process(t time.Time) {
	fmt.Printf("[%s] Payment: %s -> %.2f€ -> %s\n", t.Format(time.RFC3339), p.From, p.Amount, p.To)
}

func (p *Payment) Process() {
	now := time.Now()
	p.once.Do(func() {
		p.process(now)
	})
}

func main() {
	p := Payment{
		From:   "Jan",
		To:     "Anna",
		Amount: 123.45,
	}

	p.Process()
	p.Process()
}
