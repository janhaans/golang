package main

import (
	"fmt"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	users := []User{
		{
			FirstName: "Jan",
			LastName:  "Haans",
			Age:       54,
		},
		{
			FirstName: "Anna-Maria",
			LastName:  "Gobati",
			Age:       47,
		},
		{
			FirstName: "Valentina",
			LastName:  "Haans",
			Age:       12,
		},
	}

	fmt.Println(users)
	sort.Slice(users, func(i int, j int) bool {
		return users[i].LastName < users[j].LastName
	})
	fmt.Println(users)

}
