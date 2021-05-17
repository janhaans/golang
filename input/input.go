package main

import (
	"fmt"

	"github.com/janhaans/input/helper"
)

func main() {
	name := helper.GetName()
	fmt.Println("Hello", name)
	age := helper.GetAge()
	fmt.Printf("You are %d years old\n", age)
}
